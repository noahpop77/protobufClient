package main

import (
	"bytes"
	"context"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"testClient/olympusProto"
)

/*
Main profile:
1) /queueUp
	- PG connection for rank
	- Add party to Redis
	- Matchmaking takes place to find team mates
	- Game found and a matchID is returned for player to connect to
2) /connectToMatch
	- With the returned matchID hit up /connectToMatch
	- Only players with PUUIDs that are whitelisted can connect
		- White list determined by matchmaking server
3) /matchHistory + /riotProfile
	- Simulates the post-game client experience
*/

const (
	matchmakingURL = "http://localhost:8080/queueUp"
	gameServerURL  = "http://localhost:8081/connectToMatch"
	platformURL    = "http://localhost:8082"

	// queueTimeout bounds how long a player waits for a match. It must be
	// generous (matchmaking can legitimately take a while) but it still
	// guarantees a stuck request is eventually recycled.
	queueTimeout = 10 * time.Minute
	// apiTimeout bounds the short post-match requests.
	apiTimeout = 15 * time.Second
	// A match's simulated game length is derived from its match ID, in
	// [minGameDuration, maxGameDuration]. The game server derives the same
	// value from the same ID, so both sides agree without an extra request.
	// These bounds must match the game server's constants.
	minGameDuration = 200 * time.Second
	maxGameDuration = 400 * time.Second
	// postGameBuffer gives the game server time to write the match result to
	// the database before the client asks for it.
	postGameBuffer = 15 * time.Second
	// maxResponseBytes protects against a runaway/oversized response.
	maxResponseBytes = 8 << 20
	// maxAttempts is how many times a short API call is retried on a
	// transport error before the player's cycle is aborted.
	maxAttempts = 3
)

// matchGameDuration returns a stable per-match game length in
// [minGameDuration, maxGameDuration]. It must produce exactly the same value as
// the game server's function of the same name: it is a pure function of the
// match ID, which both sides receive from the matchmaker.
func matchGameDuration(matchID string) time.Duration {
	h := fnv.New64a()
	_, _ = h.Write([]byte(matchID))

	spanSeconds := int64((maxGameDuration - minGameDuration) / time.Second)
	offsetSeconds := int64(h.Sum64() % uint64(spanSeconds+1))
	return minGameDuration + time.Duration(offsetSeconds)*time.Second
}

// sharedTransport is used by every simulated player. The old code created a
// fresh http.Client per request (and never closed the bodies), which is what
// produced "resource temporarily unavailable" once the process ran out of
// file descriptors. One transport also lets connections be reused instead of
// opening a new TCP connection per call.
var sharedTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     false,
	MaxIdleConns:          0,     // no global cap on idle connections
	MaxIdleConnsPerHost:   20000, // the default of 2 defeats pooling at this scale
	MaxConnsPerHost:       0,     // unlimited active connections
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var (
	// queueClient has no overall timeout: /queueUp is long-lived and is
	// bounded by queueTimeout via the request context instead.
	queueClient = &http.Client{Transport: sharedTransport}
	// apiClient bounds the short platform requests.
	apiClient = &http.Client{Transport: sharedTransport, Timeout: apiTimeout}
)

func QueueUpProfile(userCount int, counter int) {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	partyRequests := []*olympusProto.Players{
		{
			PlayerPuuid:       "JvU9pQmG40YxBTfNhLKCA3Ro2WzXlJqM7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdT3",
			PlayerRiotName:    "BobbyB",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        19,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL7239D2",
			QueueType:         420,
		},
	}

	generateRandomUsers(rng, &partyRequests, userCount)

	var wg sync.WaitGroup
	for i := 0; i < len(partyRequests); i++ {
		player := partyRequests[i]
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// Each goroutine gets its own RNG: the shared *rand.Rand used by
			// the original code is not safe for concurrent use.
			playerRng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(index)))

			// Small stagger so 10k players do not all open a connection in
			// the same instant.
			time.Sleep(time.Duration(playerRng.Intn(2000)) * time.Millisecond)

			for {
				if err := runPlayerCycle(playerRng, player); err != nil {
					logFailure("player %s: %v", shortPUUID(player.PlayerPuuid), err)
					// Back off before re-queueing so a failing player does
					// not spin the CPU or hammer the server.
					time.Sleep(time.Duration(RandomIntInRange(playerRng, 1, 3)) * time.Second)
					continue
				}
			}
		}(i)
	}

	// Every player loop is intentionally infinite, so this only returns if
	// the process is interrupted.
	wg.Wait()
}

// runPlayerCycle performs one full simulated player journey: queue, connect,
// play, then read match history and profile.
func runPlayerCycle(rng *rand.Rand, player *olympusProto.Players) error {
	// 1) Queue up and wait for the matchmaker to hand us a match.
	queueCtx, cancelQueue := context.WithTimeout(context.Background(), queueTimeout)
	body, err := doProto(queueCtx, queueClient, matchmakingURL, player)
	cancelQueue()
	if err != nil {
		return fmt.Errorf("queueUp: %w", err)
	}

	var match olympusProto.MatchResponse
	if err := proto.Unmarshal(body, &match); err != nil {
		return fmt.Errorf("unmarshal queueUp response (%d bytes): %w", len(body), err)
	}
	if match.MatchID == "" {
		return fmt.Errorf("queueUp returned an empty match id")
	}

	// 2) Connect to the match we were assigned.
	connection := &olympusProto.MatchConnection{
		MatchID:          match.MatchID,
		ParticipantPUUID: player.PlayerPuuid,
		RiotName:         player.PlayerRiotName,
		RiotTag:          player.PlayerRiotTagLine,
		Rank:             strconv.Itoa(int(player.PlayerRank)),
		Role:             player.PlayerRole,
		PartyId:          player.PartyId,
		QueueType:        strconv.Itoa(int(player.QueueType)),
	}

	apiCtx, cancelAPI := context.WithTimeout(context.Background(), apiTimeout)
	_, err = doProtoRetry(apiCtx, apiClient, gameServerURL, connection)
	cancelAPI()
	if err != nil {
		return fmt.Errorf("connectToMatch: %w", err)
	}

	// 3) Simulate the player being in a game. Wait for this match's server-side
	// duration plus a buffer so the result has been written to the database.
	time.Sleep(matchGameDuration(match.MatchID) + postGameBuffer)

	// 4) Read the match history.
	historyCtx, cancelHistory := context.WithTimeout(context.Background(), apiTimeout)
	body, err = doProtoRetry(historyCtx, apiClient, platformURL+"/matchHistory",
		&olympusProto.MatchHistoryRequest{Puuid: player.PlayerPuuid})
	cancelHistory()
	if err != nil {
		return fmt.Errorf("matchHistory: %w", err)
	}
	var history olympusProto.MatchHistoryReponse
	if err := proto.Unmarshal(body, &history); err != nil {
		return fmt.Errorf("unmarshal matchHistory response (%d bytes): %w", len(body), err)
	}

	// 5) Read the summoner profile.
	profileCtx, cancelProfile := context.WithTimeout(context.Background(), apiTimeout)
	body, err = doProtoRetry(profileCtx, apiClient, platformURL+"/riotProfile",
		&olympusProto.UserProfile{Puuid: player.PlayerPuuid})
	cancelProfile()
	if err != nil {
		return fmt.Errorf("riotProfile: %w", err)
	}
	var profile olympusProto.UserProfile
	if err := proto.Unmarshal(body, &profile); err != nil {
		return fmt.Errorf("unmarshal riotProfile response (%d bytes): %w", len(body), err)
	}

	return nil
}

// doProtoRetry retries a short request on transient transport errors.
func doProtoRetry(ctx context.Context, client *http.Client, url string, msg proto.Message) ([]byte, error) {
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		body, err := doProto(ctx, client, url, msg)
		if err == nil {
			return body, nil
		}
		lastErr = err

		// A cancelled/expired context will never succeed on retry.
		if ctx.Err() != nil {
			break
		}
		if attempt < maxAttempts {
			select {
			case <-ctx.Done():
				return nil, lastErr
			case <-time.After(time.Duration(attempt) * 250 * time.Millisecond):
			}
		}
	}
	return nil, lastErr
}

// doProto marshals msg, POSTs it, and always fully reads and closes the
// response body. Reading the whole body (instead of the old chunked
// Read/Unmarshal loop) is what makes this correct for a message that spans
// multiple TCP segments: the server ends the body when the handler returns.
func doProto(ctx context.Context, client *http.Client, url string, msg proto.Message) ([]byte, error) {
	payload, err := proto.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-protobuf")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Closing here (per call) is the fix for the old `defer Body.Close()`
	// inside an infinite loop, which leaked a descriptor per iteration.
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes))
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned %d: %s", url, resp.StatusCode, truncate(string(body), 200))
	}
	if len(body) == 0 {
		return nil, fmt.Errorf("%s returned an empty body", url)
	}

	return body, nil
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

func shortPUUID(puuid string) string {
	if len(puuid) <= 10 {
		return puuid
	}
	return puuid[:10]
}

// logFailure rate-limits error logging. With 10k players a single outage can
// otherwise produce tens of thousands of identical lines per second.
var (
	logMu      sync.Mutex
	lastLog    time.Time
	suppressed int
)

func logFailure(format string, args ...any) {
	logMu.Lock()
	defer logMu.Unlock()

	if time.Since(lastLog) < time.Second {
		suppressed++
		return
	}

	msg := fmt.Sprintf(format, args...)
	if suppressed > 0 {
		log.Printf("(%d similar messages suppressed) %s", suppressed, msg)
		suppressed = 0
	} else {
		log.Printf("%s", msg)
	}
	lastLog = time.Now()
}

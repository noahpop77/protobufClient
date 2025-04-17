package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"testClient/olympusProto"
	"time"

	"google.golang.org/protobuf/proto"
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
	-
*/

func QueueUpProfile(userCount int, counter int) {
	var printMutex sync.Mutex

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	partyRequests := []*olympusProto.Players{
		{
			PlayerPuuid:       "JvU9pQmG-0YxBTfNhLKCA3Ro2WzXlJqM7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdT3",
			PlayerRiotName:    "BobbyB",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        19,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL7239D2",
			QueueType:         420,
		},
	}

	generateRandomUsers(rng, &partyRequests, userCount)

	for i := 0; i < len(partyRequests); i++ {

		go func(pr *olympusProto.Players) {

			defer func() {
				counter = counter - 1
				// atomic.AddInt64(&iterations, 1)
				// iterations += 1
			}()

			for {

				time.Sleep(time.Duration(RandomIntInRange(rng, 1, 30)) * time.Second)

				data, err := proto.Marshal(pr)
				if err != nil {
					log.Printf("Failed to marshal Protobuf: %v", err)
					return
				}

				resp, err := http.NewRequest("POST", "http://localhost:8080/queueUp", bytes.NewReader(data))
				if err != nil {
					log.Printf("Failed to send request 1: %v", err)
					return
				}
				resp.Header.Set("Content-Type", "application/x-protobuf")

				client := &http.Client{}
				streamResponse, err := client.Do(resp)
				if err != nil {
					// log.Printf("Failed to send request 2: %v", err)
					// return
					continue
				}
				defer streamResponse.Body.Close()

				var response olympusProto.MatchResponse

				// Stream the response in real-time
				scanner := bufio.NewScanner(streamResponse.Body)
				const maxBufferSize = 1024 * 1024 // Adjust this size as needed
				buf := make([]byte, maxBufferSize)
				scanner.Buffer(buf, maxBufferSize)
				for {
					// Read a chunk of data into the buffer
					n, err := streamResponse.Body.Read(buf)
					if err != nil && err != io.EOF {
						fmt.Printf("Error reading from stream: %v\n", err)
						return
					}

					// Process the data in the buffer, in chunks of Protocol Buffers messages
					data := buf[:n]
					for len(data) > 0 {
						err := proto.Unmarshal(data, &response)
						if err != nil {
							fmt.Printf("Failed to unmarshal data: %v - %s\n", err, data)
							return
						}
						// fmt.Printf("Added %s to queue...\n", pr.PlayerRiotName)
						data = data[len(data):]
					}

					// If we reached EOF, break out of the loop.
					if err == io.EOF {
						break
					}
				}

				if err := scanner.Err(); err != nil {
					log.Printf("Error reading response: %v", err)
				}

				// log.Printf("Sending...\n")

				////////////////////////////////////////////////////////////////
				// CONNECTING TO MATCH AFTER MATCHMAKING IS DONE
				////////////////////////////////////////////////////////////////

				connectionStruct := olympusProto.MatchConnection{
					MatchID:          response.MatchID, // ID of match that we got from the queueUp API endpoint
					ParticipantPUUID: pr.PlayerPuuid,
					RiotName:         pr.PlayerRiotName,
					RiotTag:          pr.PlayerRiotTagLine,
					Rank:             strconv.Itoa(int(pr.PlayerRank)),
					Role:             pr.PlayerRole,
					PartyId:          pr.PartyId,
					QueueType:        strconv.Itoa(int(pr.QueueType)),
				}

				connectionData, err := proto.Marshal(&connectionStruct)
				if err != nil {
					log.Printf("Failed to marshal result Protobuf: %v", err)
					return
				}

				connectResponse, err := http.NewRequest("POST", "http://localhost:8081/connectToMatch", bytes.NewReader(connectionData))
				if err != nil {
					log.Printf("Failed to send request 3: %v", err)
					return
				}
				connectResponse.Header.Set("Content-Type", "application/x-protobuf")

				matchClient := &http.Client{}
				streamConnectResponse, err := matchClient.Do(connectResponse)
				if err != nil {
					log.Printf("Failed to send request 4: %v, %s", err, pr)
					// return
				}
				defer streamConnectResponse.Body.Close()

				// Stream the response in real-time
				matchScanner := bufio.NewScanner(streamConnectResponse.Body)
				const maxMatchBufferSize = 1024 * 1024 // Adjust this size as needed
				matchBuf := make([]byte, maxMatchBufferSize)
				matchScanner.Buffer(matchBuf, maxBufferSize)

				// for {
				// 	// Read a chunk of data into the buffer
				// 	n, err := streamConnectResponse.Body.Read(matchBuf)
				// 	if err != nil && err != io.EOF {
				// 		fmt.Printf("Error reading from stream: %v\n", err)
				// 		return
				// 	}

				// 	// Process the data in the buffer, in chunks of Protocol Buffers messages
				// 	data := matchBuf[:n]

				// 	for len(data) > 0 {
				// 		printMutex.Lock()
				// 		fmt.Printf("%s\n", data)
				// 		printMutex.Unlock()
				// 		data = data[len(data):]
				// 	}

				// 	if err == io.EOF {
				// 		break
				// 	}
				// }

				// Simulated delay in how player behaves
				// Sleeps for a random value between 10 seconds and 1 minutes
				time.Sleep(time.Duration(RandomIntInRange(rng, 10, 60)) * time.Second)

				////////////////////////////////////////////////////////////////
				// CONNECTING TO MATCH AFTER MATCHMAKING IS DONE
				////////////////////////////////////////////////////////////////

				// Checking match history after the game is completed

				MatchHistoryRequest := olympusProto.MatchHistoryRequest{
					Puuid: pr.PlayerPuuid,
				}

				MatchHistoryRequestData, err := proto.Marshal(&MatchHistoryRequest)
				if err != nil {
					log.Printf("Failed to marshal result Protobuf: %v", err)
					return
				}

				historyResponse, err := http.NewRequest("POST", "http://localhost:8082/matchHistory", bytes.NewReader(MatchHistoryRequestData))
				if err != nil {
					log.Printf("Failed to send request 5: %v", err)
					return
				}
				historyResponse.Header.Set("Content-Type", "application/x-protobuf")

				matchHistoryClient := &http.Client{}
				matchHistoryResponse, err := matchHistoryClient.Do(historyResponse)
				if err != nil {
					log.Printf("Failed to send request 6: %v", err)
					return
				}
				defer matchHistoryResponse.Body.Close()

				var protoResponse olympusProto.MatchHistoryReponse
				body, err := io.ReadAll(matchHistoryResponse.Body)
				if err != nil {
					log.Fatalf("Failed to read response body: %v\n", err)
				}
				err = proto.Unmarshal(body, &protoResponse)
				if err != nil {
					log.Fatalf("Failed to unmarshal match history response: %v\n", err)
				}

				// Simulated delay in how player behaves
				// Sleeps for a random value between 10 seconds and 5 minutes
				time.Sleep(time.Duration(RandomIntInRange(rng, 10, 30)) * time.Second)

				////////////////////////////////////////////////////////////////
				// CONNECTING TO MATCH AFTER MATCHMAKING IS DONE
				////////////////////////////////////////////////////////////////

				// Checkign summoner profile after match history

				summonerProfileRequestBody := olympusProto.UserProfile{
					Puuid: pr.PlayerPuuid,
				}

				summonerProfileRequestData, err := proto.Marshal(&summonerProfileRequestBody)
				if err != nil {
					log.Printf("Failed to marshal result Protobuf: %v", err)
					return
				}

				summonerProfileRequest, err := http.NewRequest("POST", "http://localhost:8082/riotProfile", bytes.NewReader(summonerProfileRequestData))
				if err != nil {
					log.Printf("Failed to send request 7: %v", err)
					return
				}
				historyResponse.Header.Set("Content-Type", "application/x-protobuf")

				summonerProfileClient := &http.Client{}
				summonerProfileResponse, err := summonerProfileClient.Do(summonerProfileRequest)
				if err != nil {
					log.Printf("Failed to send request 8: %v", err)
					return
				}
				defer matchHistoryResponse.Body.Close()

				var summonerProfileProtoResponse olympusProto.UserProfile
				profileBody, err := io.ReadAll(summonerProfileResponse.Body)
				if err != nil {
					log.Fatalf("Failed to read response body: %v\n", err)
				}
				err = proto.Unmarshal(profileBody, &summonerProfileProtoResponse)
				if err != nil {
					// log.Fatalf("Failed to unmarshal summoner profile response: %v\n", err)
					log.Printf("Failed to unmarshal summoner profile response: %v\n", err) // ProfileBody shows "No Data Found"
					return                                                                 // Exit the goroutine
				}

				// fmt.Printf("Player data: %s(trunc), %s, %s, %d, %d, %d\n", summonerProfileProtoResponse.Puuid[0:10], summonerProfileProtoResponse.RiotName, summonerProfileProtoResponse.RiotTag, summonerProfileProtoResponse.Rank, summonerProfileProtoResponse.Wins, summonerProfileProtoResponse.Losses)

				// Simulated delay in how player behaves
				// Sleeps for a random value between 10 seconds and 3 minutes
				time.Sleep(time.Duration(RandomIntInRange(rng, 10, 18)) * time.Second)

				// log.Printf("Rotation Complete")
				printMutex.Lock()

				// atomic.AddInt64(&iterations, 1)
				log.Printf("counter: %d", counter)
				printMutex.Unlock()
			}

		}(partyRequests[i])
	}
	select {}

}

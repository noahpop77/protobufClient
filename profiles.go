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
	// var printMutex sync.Mutex

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

	// fmt.Printf("Pre GenerateRandomUsers\n")
	generateRandomUsers(rng, &partyRequests, userCount)
	var counterMutex sync.Mutex
	// fmt.Printf("Pre looping to create all the go func go-routines\n")
	for i := 0; i < len(partyRequests); i++ {

		go func(pr *olympusProto.Players) {

			defer func() {
				counterMutex.Lock()
				counter--
				counterMutex.Unlock()
			}()

			for {

				// time.Sleep(time.Duration(RandomIntInRange(rng, 1, 5)) * time.Second)
				// fmt.Printf("Running Request Loop\n")
				data, err := proto.Marshal(pr)
				if err != nil {
					log.Printf("Failed to marshal Protobuf: %v", err)
					return
				}

				// fmt.Printf("Pre Queueing Up Request\n")
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

				/*
					Note: So we find everyone a match but error later?

					Server:
					platform_server       | 2025/12/02 04:20:41 No match found for PUUID: S4URCScuWqJQWMY22h979IVPjVTEpeY7cVxPGOOXN5xeqm8Wc1eVlItPZfGLNLK8, , , 0
					game_server           | 2025/12/02 04:20:41 end count 6
					platform_server       | 2025/12/02 04:20:41 No match found for PUUID: vmoipFWbGBA3Eqr88Evrc3QFLml3fJ6lDfviohl6rHb459VUgJj8KmQ6chveWOcD, , , 0

					Tester:
					2025/12/01 23:20:41 S4URCScuWqJQWMY22h979IVPjVTEpeY7cVxPGOOXN5xeqm8Wc1eVlItPZfGLNLK8, MATCH_5W5SX7PSYC
					2025/12/01 23:20:41 counter: 100
					2025/12/01 23:20:41 vmoipFWbGBA3Eqr88Evrc3QFLml3fJ6lDfviohl6rHb459VUgJj8KmQ6chveWOcD, MATCH_3DYFWP37YT
					2025/12/01 23:20:41 counter: 100
				*/
				// log.Printf("%v, %v, %v", pr.PlayerPuuid, response.MatchID, response.Participants)

				// // time.Sleep(time.Duration(RandomIntInRange(rng, 1, 5)) * time.Second)

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
				// fmt.Printf("Pre Connecting to Match Request\n")

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
					return
				}
				defer streamConnectResponse.Body.Close()

				// if streamConnectResponse.StatusCode != http.StatusOK {
				// 	body, _ := io.ReadAll(streamConnectResponse.Body)
				// 	log.Printf("✗ Server returned error status %d: %s", streamConnectResponse.StatusCode, body)
				// 	return
				// }
				// log.Println("✓ Server returned 200 OK")

				// // Stream the response in real-time
				// matchScanner := bufio.NewScanner(streamConnectResponse.Body)
				// const maxMatchBufferSize = 1024 * 1024 // Adjust this size as needed
				// matchBuf := make([]byte, maxMatchBufferSize)
				// matchScanner.Buffer(matchBuf, maxBufferSize)

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
				time.Sleep(time.Duration(RandomIntInRange(rng, 5, 6)) * time.Second)

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

				matchHistoryClient := &http.Client{} //2025/04/16 18:38:11 Failed to send request 6: Post "http://localhost:8082/matchHistory": read tcp [::1]:57101->[::1]:8082: read: connection reset by peer
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

				// Note: So the failed requests are failing before match history
				// log.Printf("%s\n%v", protoResponse.Puuid, protoResponse.Matches)

				// Simulated delay in how player behaves
				// Sleeps for a random value between 10 seconds and 5 minutes
				// time.Sleep(time.Duration(RandomIntInRange(rng, 11, 15)) * time.Second)

				////////////////////////////////////////////////////////////////
				// CONNECTING TO MATCH AFTER MATCHMAKING IS DONE
				////////////////////////////////////////////////////////////////

				// Checkign summoner profile after match history

				//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
				// Note: TODO: The matchmaking connection api request previously seems to sometimes prematurely exit out or die somewhere
				// on the server so it cancels out here as well and speeds to the riotprofile request. Its a big problem. FIX IT
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
				summonerProfileRequest.Header.Set("Content-Type", "application/x-protobuf")

				summonerProfileClient := &http.Client{}
				summonerProfileResponse, err := summonerProfileClient.Do(summonerProfileRequest)
				if err != nil {
					log.Printf("Failed to send request 8: %v", err)
					return
				} //2025/04/16 18:38:33 Failed to send request 8: Post "http://localhost:8082/riotProfile": dial tcp [::1]:8082: connect: resource temporarily unavailable
				defer summonerProfileResponse.Body.Close()

				var summonerProfileProtoResponse olympusProto.UserProfile
				profileBody, err := io.ReadAll(summonerProfileResponse.Body)
				if err != nil {
					log.Fatalf("Failed to read response body: %v\n", err)
				}
				err = proto.Unmarshal(profileBody, &summonerProfileProtoResponse)
				if err != nil {
					// log.Fatalf("Failed to unmarshal summoner profile response: %v\n", err)
					log.Printf("Failed to unmarshal summoner profile response: %v\n", err) // ProfileBody shows "No Data Found"
					//log.Printf("return info: %v", summonerProfileProtoResponse.RiotName)
					return // Exit the goroutine
				}
				//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

				// fmt.Printf("Player data: %s(trunc), %s, %s, %d, %d, %d\n", summonerProfileProtoResponse.Puuid[0:10], summonerProfileProtoResponse.RiotName, summonerProfileProtoResponse.RiotTag, summonerProfileProtoResponse.Rank, summonerProfileProtoResponse.Wins, summonerProfileProtoResponse.Losses)

				// Simulated delay in how player behaves
				// Sleeps for a random value between 10 seconds and 3 minutes
				// time.Sleep(time.Duration(RandomIntInRange(rng, 1, 5)) * time.Second)

				// log.Printf("Rotation Complete")

				// printMutex.Lock()
				// log.Printf("counter: %d", counter)
				// printMutex.Unlock()
			}

		}(partyRequests[i])
	}
	select {}

}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

	"google.golang.org/protobuf/proto"

	// Import the generated Protobuf Go file
	"testClient/matchmaking/party"
)

func main() {

	partyRequests := []*party.Players{
		{
			PlayerPuuid:       "Xo7QpJdK-2YgVzLtRAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWA",
			PlayerRiotName:    "Spica",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        18,
			PlayerRole:        "Bottom",
			PartyId:           "PARTY_5JKL723LJ2",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "JvU9pQmG-0YxBTfNhLKCA3Ro2WzXlJqM7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdT3",
			PlayerRiotName:    "BobbyB",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        19,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL7239D2",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9B",
			PlayerRiotName:    "Haidder",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        20,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL723MK1",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YB",
			PlayerRiotName:    "Mingle",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        21,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723VQ7",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybNX",
			PlayerRiotName:    "Dingle",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        21,
			PlayerRole:        "Jungle",
			PartyId:           "PARTY_5JKL723K8L",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "2YgVzLtRAEm6c0f5_WxA8PqBN3l9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0x",
			PlayerRiotName:    "Karar",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        22,
			PlayerRole:        "Support",
			PartyId:           "PARTY_5JKL723H4M",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3o",
			PlayerRiotName:    "Anghel",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        22,
			PlayerRole:        "Jungle",
			PartyId:           "PARTY_5JKL723FQ9",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "RAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0YxBTfNh7",
			PlayerRiotName:    "Zizzy",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        18,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723D1Y",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "WJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-BBB",
			PlayerRiotName:    "Fruity",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        20,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL723UQ3",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "F6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhLP",
			PlayerRiotName:    "Mathmood",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        19,
			PlayerRole:        "Support",
			PartyId:           "PARTY_5JKL723J7L",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "M7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0",
			PlayerRiotName:    "Zeus",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        23,
			PlayerRole:        "Bottom",
			PartyId:           "PARTY_5JKL723I9N",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybXoEWsK1L9RBpCdTJ3KCA3W",
			PlayerRiotName:    "Apollo",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        22,
			PlayerRole:        "Support",
			PartyId:           "PARTY_5JKL723Q6F",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7Pg",
			PlayerRiotName:    "Butts",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        42,
			PlayerRole:        "Jungle",
			PartyId:           "PARTY_5JKL723U8F",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "BTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvUp",
			PlayerRiotName:    "Guts",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        39,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL723V6Y",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCT",
			PlayerRiotName:    "Nuts",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        38,
			PlayerRole:        "Bottom",
			PartyId:           "PARTY_5JKL723N1Z",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "1dF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNh7",
			PlayerRiotName:    "Gaze",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        38,
			PlayerRole:        "Support",
			PartyId:           "PARTY_5JKL723V5X",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3gl41C",
			PlayerRiotName:    "Athena",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        39,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723L7Y",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "WJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4o",
			PlayerRiotName:    "Hera",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        40,
			PlayerRole:        "Jungle",
			PartyId:           "PARTY_5JKL723Q2K",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "RAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0YxBFHJBB",
			PlayerRiotName:    "Achiles",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        41,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL723D6W",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEW-",
			PlayerRiotName:    "Hercules",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        30,
			PlayerRole:        "Bottom",
			PartyId:           "PARTY_5JKL723L3T",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4X",
			PlayerRiotName:    "Dionysus",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        38,
			PlayerRole:        "Support",
			PartyId:           "PARTY_5JKL723N4C",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "JvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1",
			PlayerRiotName:    "Artemis",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        39,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723M9Y",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "M7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0Yx",
			PlayerRiotName:    "Hermes",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        40,
			PlayerRole:        "Jungle",
			PartyId:           "PARTY_5JKL723X0B",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "BTfNhL7P5g8ZVdF6tAybN4QoS1pZGXWJvU9p-GFDGSDFGD5623654GDFGHB5234BDB",
			PlayerRiotName:    "Hephaestus",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        42,
			PlayerRole:        "Top",
			PartyId:           "PARTY_5JKL723S6Q",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "2YgVzLtRAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0Yx",
			PlayerRiotName:    "Zephyrus",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        43,
			PlayerRole:        "Bottom",
			PartyId:           "PARTY_5JKL723Y5B",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6",
			PlayerRiotName:    "Ares",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        23,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723B8F",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3Xo",
			PlayerRiotName:    "Poseidon",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        20,
			PlayerRole:        "Jungle",
			PartyId:           "PARTY_5JKL723B1M",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "BTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9p",
			PlayerRiotName:    "Hades",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        20,
			PlayerRole:        "Bottom",
			PartyId:           "PARTY_5JKL723R4M",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoW",
			PlayerRiotName:    "Shmekkles",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        40,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723G9W",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBf",
			PlayerRiotName:    "NMOAF",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        41,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723D8X",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "",
			PlayerRiotName:    "",
			PlayerRiotTagLine: "",
			PlayerRank:        0,
			PlayerRole:        "",
			PartyId:            "",
			QueueType:          6969,
		},
		{
			PlayerPuuid:       "TotallyWrongID",
			PlayerRiotName:    "ErrorUser1",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        40,
			PlayerRole:        "Middle",
			PartyId:           "PARTY_5JKL723D81",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "RRZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxRR",
			PlayerRiotName:    "ErrorUser3",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        40,
			PlayerRole:        "Jungle",
			PartyId:           "PARTY_5JKL723D82",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "PPZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxOO",
			PlayerRiotName:    "ErrorUser5",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        40,
			PlayerRole:        "Support",
			PartyId:           "PARTY_5JKL723D83",
			QueueType:         420,
		},
		{
			PlayerPuuid:       "HHZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxAA",
			PlayerRiotName:    "ErrorUser7",
			PlayerRiotTagLine: "NA1",
			PlayerRank:        40,
			PlayerRole:        "Bottom",
			PartyId:           "PARTY_5JKL723D84",
			QueueType:         420,
		},
	}

	var wg sync.WaitGroup

	for i := 0; i < len(partyRequests); i++ {
		wg.Add(1)
		go func(pr *party.Players) {
			defer wg.Done()

			data, err := proto.Marshal(pr)
			if err != nil {
				log.Printf("Failed to marshal Protobuf: %v", err)
				return
			}

			resp, err := http.NewRequest("POST", "http://localhost:8080/queueUp", bytes.NewReader(data))
			if err != nil {
				log.Printf("Failed to send request: %v", err)
				return
			}
			resp.Header.Set("Content-Type", "application/x-protobuf")

			client := &http.Client{}
			streamResponse, err := client.Do(resp)
			if err != nil {
				log.Printf("Failed to send request: %v", err)
				return
			}
			defer streamResponse.Body.Close()

			var response party.MatchResponse

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
					fmt.Printf("Added %s to queue...\n", pr.PlayerRiotName)
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


			

			connectionStruct := party.MatchConnection{
				MatchID: response.MatchID,
				ParticipantPUUID: pr.PlayerPuuid,
				RiotName: pr.PlayerRiotName,
				RiotTag: pr.PlayerRiotTagLine,
				Rank: strconv.Itoa(int(pr.PlayerRank)),
				Role: pr.PlayerRole,
				PartyId: pr.PartyId,
				QueueType: strconv.Itoa(int(pr.QueueType)),
			}

			connectionData, err := proto.Marshal(&connectionStruct)
			if err != nil {
				log.Printf("Failed to marshal result Protobuf: %v", err)
				return
			}
			
			connectResponse, err := http.NewRequest("POST", "http://localhost:8081/connectToMatch", bytes.NewReader(connectionData))
			if err != nil {
				log.Printf("Failed to send request: %v", err)
				return
			}
			connectResponse.Header.Set("Content-Type", "application/x-protobuf")

			matchClient := &http.Client{}
			streamConnectResponse, err := matchClient.Do(connectResponse)
			if err != nil {
				log.Printf("Failed to send request: %v", err)
				return
			}
			defer streamConnectResponse.Body.Close()

			
			var printMutex sync.Mutex
			// Stream the response in real-time
			matchScanner := bufio.NewScanner(streamConnectResponse.Body)
			const maxMatchBufferSize = 1024 * 1024 // Adjust this size as needed
			matchBuf := make([]byte, maxMatchBufferSize)
			matchScanner.Buffer(matchBuf, maxBufferSize)
			//connectionResultStruct := party.MatchResult{}

			for {
				// Read a chunk of data into the buffer
				n, err := streamConnectResponse.Body.Read(matchBuf)
				if err != nil && err != io.EOF {
					fmt.Printf("Error reading from stream: %v\n", err)
					return
				}
			
				// Process the data in the buffer, in chunks of Protocol Buffers messages
				data := matchBuf[:n]
				
				for len(data) > 0 {
					printMutex.Lock()
					fmt.Printf("%s\n", data)
					printMutex.Unlock()
					data = data[len(data):]
				}

				if err == io.EOF {
					break
				}
			}



			MatchHistoryRequest := party.MatchHistoryRequest{
				Puuid: pr.PlayerPuuid,
			}

			MatchHistoryRequestData, err := proto.Marshal(&MatchHistoryRequest)
			if err != nil {
				log.Printf("Failed to marshal result Protobuf: %v", err)
				return
			}

			historyResponse, err := http.NewRequest("POST", "http://localhost:8082/matchHistory", bytes.NewReader(MatchHistoryRequestData))
			if err != nil {
				log.Printf("Failed to send request: %v", err)
				return
			}
			historyResponse.Header.Set("Content-Type", "application/x-protobuf")

			matchHistoryClient := &http.Client{}
			matchHistoryResponse, err := matchHistoryClient.Do(historyResponse)
			if err != nil {
				log.Printf("Failed to send request: %v", err)
				return
			}
			defer matchHistoryResponse.Body.Close()
			
			var protoResponse party.MatchHistoryReponse
			body, err := io.ReadAll(matchHistoryResponse.Body)
			if err != nil {
				log.Fatalf("Failed to read response body: %v\n", err)
			}
			err = proto.Unmarshal(body, &protoResponse)
			if err != nil {
				log.Fatalf("Failed to unmarshal match history response: %v\n", err)
			}

			// fmt.Printf("MatchID: %s\n", protoResponse.Matches[0].MatchID)

			fmt.Printf("%s\n", protoResponse.Puuid)
			// fmt.Printf("%s\n", protoResponse.Matches[0])
			// fmt.Printf("History: %s\n", protoResponse.Matches)
			for _, value := range protoResponse.Matches{
				fmt.Printf("MatchID: %s\nGameDuration: %s\n", value.MatchID, value.GameDuration)
			}

		}(partyRequests[i])
	}

	wg.Wait()

}

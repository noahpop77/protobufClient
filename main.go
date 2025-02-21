package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"

	// Import the generated Protobuf Go file
	"testClient/matchmaking/party"
)

func main() {

	partyRequests := []*party.PartyRequest{
		{
			PartyId:  "PARTY_5JKL723LJ2",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Spica",
					RiotTagLine: "NA1",
					Rank:        "18",
					Role:        "Bottom",
					Puuid:       "Xo7QpJdK-2YgVzLtRAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWA",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL7239D2",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "BobbyB",
					RiotTagLine: "NA1",
					Rank:        "19",
					Role:        "Top",
					Puuid:       "JvU9pQmG-0YxBTfNhLKCA3Ro2WzXlJqM7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdT3",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723MK1",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Haidder",
					RiotTagLine: "NA1",
					Rank:        "20",
					Role:        "Top",
					Puuid:       "KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9B",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723VQ7",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Mingle",
					RiotTagLine: "NA1",
					Rank:        "21",
					Role:        "Middle",
					Puuid:       "P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YB",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723K8L",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Dingle",
					RiotTagLine: "NA1",
					Rank:        "21",
					Role:        "Jungle",
					Puuid:       "L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybNX",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723H4M",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Karar",
					RiotTagLine: "NA1",
					Rank:        "22",
					Role:        "Support",
					Puuid:       "2YgVzLtRAEm6c0f5_WxA8PqBN3l9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0x",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723FQ9",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Anghel",
					RiotTagLine: "NA1",
					Rank:        "22",
					Role:        "Jungle",
					Puuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3o",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723D1Y",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Zizzy",
					RiotTagLine: "NA1",
					Rank:        "18",
					Role:        "Middle",
					Puuid:       "RAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0YxBTfNh7",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723UQ3",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Fruity",
					RiotTagLine: "NA1",
					Rank:        "20",
					Role:        "Top",
					Puuid:       "WJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4o",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723J7L",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Mathmood",
					RiotTagLine: "NA1",
					Rank:        "19",
					Role:        "Support",
					Puuid:       "F6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhLP",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723I9N",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Zeus",
					RiotTagLine: "NA1",
					Rank:        "23",
					Role:        "Bottom",
					Puuid:       "M7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723Q6F",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Apollo",
					RiotTagLine: "NA1",
					Rank:        "22",
					Role:        "Support",
					Puuid:       "N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybXoEWsK1L9RBpCdTJ3KCA3W",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723B8F",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Ares",
					RiotTagLine: "NA1",
					Rank:        "23",
					Role:        "Middle",
					Puuid:       "XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723B1M",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Poseidon",
					RiotTagLine: "NA1",
					Rank:        "20",
					Role:        "Jungle",
					Puuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3Xo",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723R4M",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Hades",
					RiotTagLine: "NA1",
					Rank:        "20",
					Role:        "Bottom",
					Puuid:       "BTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9p",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723G9W",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Shmekkles",
					RiotTagLine: "NA1",
					Rank:        "40",
					Role:        "Middle",
					Puuid:       "yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoW",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723D8X",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "NMOAF",
					RiotTagLine: "NA1",
					Rank:        "41",
					Role:        "Middle",
					Puuid:       "g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBf",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723U8F",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Butts",
					RiotTagLine: "NA1",
					Rank:        "42",
					Role:        "Jungle",
					Puuid:       "tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7Pg",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723V6Y",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Guts",
					RiotTagLine: "NA1",
					Rank:        "39",
					Role:        "Top",
					Puuid:       "BTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvUp",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723N1Z",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Nuts",
					RiotTagLine: "NA1",
					Rank:        "38",
					Role:        "Bottom",
					Puuid:       "XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCT",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723V5X",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Gaze",
					RiotTagLine: "NA1",
					Rank:        "38",
					Role:        "Support",
					Puuid:       "1dF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNh7",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723L7Y",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Athena",
					RiotTagLine: "NA1",
					Rank:        "39",
					Role:        "Middle",
					Puuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3o",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723Q2K",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Hera",
					RiotTagLine: "NA1",
					Rank:        "40",
					Role:        "Jungle",
					Puuid:       "WJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4o",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723D6W",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Achiles",
					RiotTagLine: "NA1",
					Rank:        "41",
					Role:        "Top",
					Puuid:       "RAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0YxBTfNh7",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723L3T",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Hercules",
					RiotTagLine: "NA1",
					Rank:        "30",
					Role:        "Bottom",
					Puuid:       "pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEW-",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723N4C",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Dionysus",
					RiotTagLine: "NA1",
					Rank:        "38",
					Role:        "Support",
					Puuid:       "L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4X",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723M9Y",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Artemis",
					RiotTagLine: "NA1",
					Rank:        "39",
					Role:        "Middle",
					Puuid:       "JvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723X0B",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Hermes",
					RiotTagLine: "NA1",
					Rank:        "40",
					Role:        "Jungle",
					Puuid:       "M7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0Yx",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723S6Q",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Hephaestus",
					RiotTagLine: "NA1",
					Rank:        "42",
					Role:        "Top",
					Puuid:       "QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybN4XoEWsK1L9RBpCdTJ3KCA3-",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723Y5B",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Zephyrus",
					RiotTagLine: "NA1",
					Rank:        "43",
					Role:        "Bottom",
					Puuid:       "2YgVzLtRAEm6c0f5_WxA8PqBN3yL9vMGbJhFzsTUKCdN4QoS1pZGXWJvU9pQmG0Yx",
				},
			},
		},
		{
			PartyId:  "PARTY_5JKL723B2F",
			TeamCount: "1",
			QueueType: 420,
			Participants: []*party.Participant{
				{
					RiotName:    "Hera",
					RiotTagLine: "NA1",
					Rank:        "30",
					Role:        "Support",
					Puuid:       "N9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybNA",
				},
			},
		},
	}

	// Flat datastructure for people who queue up
	// - Filter for puuid 
	// - No need for 2 entries per person
	// {
	// 	Player1Puuid:       "N9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybNA",
	// 	Player1PartyId:  	"PARTY_5JKL723B2F",
	// 	Player1RiotName:    "Hera",
	// 	Player1RiotTagLine: "NA1",
	// 	Player1Rank:        "30",
	// 	Player1Role:        "",
	// 	Player2Puuid:       "",
	// 	Player2PartyId:  	"",
	// 	Player2RiotName:    "",
	// 	Player2RiotTagLine: "",
	// 	Player2Rank:        "",
	// 	Player2Role:        "",
	// 	TeamCount: 	 "1",
	// 	QueueType: 	 420,
	// },
	

	// fmt.Printf("PartyRequest: %+v\n", partyRequest)

	targetUser := 20
	
	// for i := 0; i < len(partyRequests); i++ {
	for i := targetUser; i < (targetUser + 1); i++ {
		// Serialize the Protobuf message to binary format
		data, err := proto.Marshal(partyRequests[i])
		if err != nil {
			log.Fatalf("Failed to marshal Protobuf: %v", err)
		}

		// Send the request
		start := time.Now()
		resp, err := http.Post("http://localhost:8080/queueUp", "application/x-protobuf", bytes.NewReader(data))
		elapsed := time.Since(start).Microseconds()
		if err != nil {
			log.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		responseBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("%s: Request took(μs) %d\n%s", resp.Status, elapsed, responseBody)



		// Send the request
		startMatchmaking := time.Now()
		matchmakingResponse, err := http.Post("http://localhost:8080/matchmaking", "application/x-protobuf", bytes.NewReader(data))
		elapsedMatchmaking := time.Since(startMatchmaking).Microseconds()
		if err != nil {
			log.Fatalf("Failed to send request: %v", err)
		}
		defer matchmakingResponse.Body.Close()

		parsedResponse, _ := io.ReadAll(matchmakingResponse.Body)
		fmt.Printf("%s: Request took(μs) %d\n%s", matchmakingResponse.Status, elapsedMatchmaking, parsedResponse)
	}
	
}
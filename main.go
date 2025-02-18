package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/protobuf/proto"

	// Import the generated Protobuf Go file
	"testClient/matchmaking/party" // Adjust the path to your generated pb.go file
)

func main() {
	// partyRequest := &party.PartyRequest{
	// 	PartyId:  "PARTY_5JKL723LJ1",
	// 	TeamCount: "2",
	// 	QueueType: 420,
	// 	Participants: []*party.Participant{
	// 		{
	// 			RiotName:    "bsawatestuser",
	// 			RiotTagLine: "NA1",
	// 			Rank: "32",
	// 			Puuid:       "HgWwc6_3QYsDbqi4TmEMdhscy7MDTAIr2iBEoh8Nn-HkqHhe1PZDh442GuUNa6ipV8dqgJUNN2KlpQ",
	// 		},
	// 	},
	// }

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
					Puuid:       "N9RBpCdTJ3KCA3XoEWs-N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybNA",
				},
			},
		},
	}
	

	// fmt.Printf("PartyRequest: %+v\n", partyRequest)
	for i := 0; i < len(partyRequests); i++ {
		// Serialize the Protobuf message to binary format
		data, err := proto.Marshal(partyRequests[i])
		if err != nil {
			log.Fatalf("Failed to marshal Protobuf: %v", err)
		}

		// Send the request
		resp, err := http.Post("http://localhost:8080/matchmaking", "application/x-protobuf", bytes.NewReader(data))
		if err != nil {
			log.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		fmt.Println("Response Status:", resp.Status)
	}
	
}
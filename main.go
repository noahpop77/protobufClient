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
	partyRequest := &party.PartyRequest{
		PartyId:  "PARTY_5JKL723LJ1",
		TeamCount: "2",
		QueueType: 420,
		Participants: []*party.Participant{
			{
				RiotName:    "bsawatestuser",
				RiotTagLine: "NA1",
				Rank: "32",
				Puuid:       "HgWwc6_3QYsDbqi4TmEMdhscy7MDTAIr2iBEoh8Nn-HkqHhe1PZDh442GuUNa6ipV8dqgJUNN2KlpQ",
			},
		},
	}

	// fmt.Printf("PartyRequest: %+v\n", partyRequest)

	// Serialize the Protobuf message to binary format
	data, err := proto.Marshal(partyRequest)
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
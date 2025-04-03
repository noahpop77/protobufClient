package main

import (
	"fmt"
	"log"
	"math/rand"
	"testClient/olympusProto"
)

func generateRandomUsers(rng *rand.Rand, partyRequests *[]*olympusProto.Players, userCount int){
	log.Printf("In QueueUpProfile")
	for i := 0; i < userCount; i++ {
		*partyRequests = append(*partyRequests, &olympusProto.Players{
			PlayerPuuid:       RandomString(rng, 64),
			PlayerRiotName:    RandomString(rng, RandomIntInRange(rng, 6, 30)),
			PlayerRiotTagLine: RandomString(rng, RandomIntInRange(rng, 3, 5)),
			PlayerRank:        int32(RandomIntInRange(rng, 1, 44)),
			PlayerRole:        RandomRole(rng),
			PartyId:           fmt.Sprintf("PARTY_%s", RandomString(rng, 10)),
			QueueType:         420,
		})
		log.Printf("Appending")
	}

}

func RandomRole(rng *rand.Rand) string{
	roles := []string{
		"Middle",
		"Top",
		"Support",
		"Jungle",
		"Bottom",
	}
	return roles[RandomIntInRange(rng, 0, 4)]
}

func RandomIntInRange(rng *rand.Rand, min, max int) int {
	if min > max {
		log.Fatalf("min cannot be greater than max")
	}
	return rng.Intn(max-min+1) + min
}

// Based off of the letters string we can generate random strings
// based off a provided length and the characters provided
func RandomString(rng *rand.Rand, length int) string {
	const letters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rng.Intn(len(letters))]
	}
	return string(result)
}

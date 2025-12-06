package main

import (
	"flag"
	"fmt"
)

/*
{
	PlayerPuuid:       "N4QoS1pZGXWJvU9pQmG0YxBTfNhL7P5g8ZVdF6tAybXoEWsK1L9RBpCdTJ3KCA3W",
	PlayerRiotName:    "Apollo",
	PlayerRiotTagLine: "NA1",
	PlayerRank:        22,
	PlayerRole:        "Support",
	PartyId:           "PARTY_5JKL723Q6F",
	QueueType:         420,
},
*/

func main() {
	printBanner()

	// Define a flag
	echo := flag.String("echo", "default", "Echos your argument to the screen")
	userCount := flag.Int("users", 100, "Number of simulated users")

	// Parse the flags
	flag.Parse()

	// Print the value of the flag
	fmt.Printf("Flags used:\nName: %s\nUser Count: %d\n", *echo, *userCount)
	counter := *userCount
	QueueUpProfile(*userCount, counter)
}

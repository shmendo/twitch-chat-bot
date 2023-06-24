package main

import (
	"log"
	"os"
	"sync"

	"github.com/shmendo/twitch-chat-bot/pkg/irc"
)

func main() {
	var (
		endpoint = os.Getenv("ENDPOINT")
		username = os.Getenv("USERNAME")
		password = os.Getenv("API_TOKEN")
		channel  = os.Getenv("CHANNEL")
		wg       sync.WaitGroup
	)

	client, err := irc.NewClient(endpoint)
	if err != nil {
		log.Println("Error instantiating IRC Client", err)
		os.Exit(1)
	}

	// First, let's set up our message handler
	client.OnMessage(HandleMessage)

	// listen forever in the background
	wg.Add(1)
	go client.ListenForMessages()

	// Then, let's authenticate
	err = client.Authenticate(username, password)
	if err != nil {
		log.Println("could not authenticate!", err)
		os.Exit(1)
	}

	// Finally, let's join a room
	err = client.JoinChannel(channel)
	if err != nil {
		log.Printf("could not join %s!", channel)
		log.Println(err)
		os.Exit(1)
	}

	wg.Wait()
	log.Println("Exiting")
}

func HandleMessage(message string) {
	log.Println("OnMessage()", message)

}

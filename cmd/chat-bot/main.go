package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/shmendo/twitch-chat-bot/pkg/chuck"
	"github.com/shmendo/twitch-chat-bot/pkg/irc"
)

func main() {
	var (
		endpoint = os.Getenv("ENDPOINT")
		username = os.Getenv("USERNAME")
		password = os.Getenv("API_TOKEN")
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
	client.Authenticate(username, password)

	// Wait forever :)
	wg.Wait()
}

// The Meat & Potatoes
func HandleMessage(message irc.Message, replyWith irc.ReplyCallback) {
	if message.CommandType == "bot" && message.Command.Command == "chucknorris" {
		fact, err := chuck.RandomChuckFact()
		if err != nil {
			log.Println("Could not retrieve chucknorris fact, he may have roundhouse kicked the server!", err)
		}
		log.Println(fact)
		log.Println(message.Source.Nick)
		log.Println(message.Source.Host)
		log.Println(message.Command.Command)
		log.Println(message.Command.Channel)
		log.Println(message.Command.Info)
		log.Println(message.Parameters)
		// replyWith()
	} else {
		switch message.Command.Command {
		case "JOIN":
			log.Printf("Successfully joined %s", message.Command.Channel)
			break
		case "376":
			// We have completed auth, let's join the room
			channelName := os.Getenv("CHANNEL")
			replyWith(fmt.Sprintf("JOIN #%s", channelName))
			break
		default:
			// log.Printf(
			// 	"Nothing to do for - Command: %s, Channel: %s, Info: %s",
			// 	message.Command.Command,
			// 	message.Command.Channel,
			// 	message.Command.Info,
			// )
		}
	}
}

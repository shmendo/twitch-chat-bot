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
	if message.CommandType == "bot" && message.Parameters == "!chucknorris" {
		fact, err := chuck.RandomChuckFact()
		if err != nil {
			log.Println("Could not retrieve chucknorris fact, he may have roundhouse kicked the server!", err)
		}
		replyWith(fmt.Sprintf("PRIVMSG %s :%s", message.Command.Channel, fact))
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
			// 	"UNRECOGNIZED - Command: %s, Channel: %s, Info: %s, Parameters: %s",
			// 	message.Command.Command,
			// 	message.Command.Channel,
			// 	message.Command.Info,
			// 	message.Parameters,
			// )
		}
	}
}

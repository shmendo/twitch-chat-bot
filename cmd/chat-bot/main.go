package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

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

	// General message handler for ANY message
	client.OnMessage(HandleMessage)
	client.OnAuthenticated(HandleAuthenticated)
	client.OnPrivateMessage(HandlePrivateMessage)
	client.OnBotCommand(HandleBotCommand)

	// listen forever in the background
	wg.Add(1)
	go client.ListenForMessages()
	client.Authenticate(username, password)

	// Wait forever :)
	wg.Wait()
}

func Log(message string) {
	timestamp := time.Now()
	log.Printf("[%s] %s", timestamp, message)
}

func HandleAuthenticated(message irc.Message, replyWith irc.ReplyCallback) {
	log.Println("TwitchChatBot->HandleAuthenticated()", message, log.Lshortfile)
	channelName := os.Getenv("CHANNEL")
	replyWith(fmt.Sprintf("JOIN #%s", channelName))
}

func HandlePrivateMessage(message irc.Message, replyWith irc.ReplyCallback) {
	log.Println("TwitchChatBot->HandlePrivateMessage()", message, log.Lshortfile)
}

func HandleBotCommand(message irc.Message, replyWith irc.ReplyCallback) {
	log.Println("TwitchChatBot->HandleBotCommand()", message, log.Lshortfile)
	if message.Parameters == "!chucknorris" {
		fact, err := chuck.RandomChuckFact()
		if err != nil {
			log.Println("Could not retrieve chucknorris fact, he may have roundhouse kicked the server!")
		}
		replyWith(fmt.Sprintf("PRIVMSG %s :%s", message.Command.Channel, fact))
	} else {
		log.Println("not handling bot command %s", message.Parameters)
	}
}

// The Meat & Potatoes
func HandleMessage(message irc.Message, replyWith irc.ReplyCallback) {
	// log.Println("TwitchChatBot->HandleMessage()", message, log.Lshortfile)
}

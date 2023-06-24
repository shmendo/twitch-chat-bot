package main

import (
	"fmt"
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
	log.Println("Exiting")
}

func HandleMessage(message irc.Message, replyWith irc.ReplyCallback) {
	switch message.Command.Command {
	case "JOIN":
		log.Printf("Successfully joined %s", message.Command.Channel)
		break
	case "PART":
		break
	case "NOTICE":
		break
	case "CLEARCHAT":
		break
	case "HOSTTARGET":
		break
	case "PRIVMSG":
		break
	case "PING":
		break
	case "CAP":
		break
	case "GLOBALUSERSTATE":
		break
	case "USERSTATE":
		break
	case "ROOMSTATE":
		break
	case "RECONNECT":
		break
	case "421":
		break
	case "001":
		replyWith(fmt.Sprintf("JOIN %s", os.Getenv("channel")))
		break
	case "002":
		break
	case "003":
		break
	case "004":
		break
	case "353":
		break
	case "366":
		break
	case "372":
		break
	case "375":
		break
	case "376":
		break
	default:
		log.Printf("UNKNOWN: %s", message.Text)
	}
}

func Authenticate() {

}

package irc

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type MessageHandler func(message string)

type client struct {
	conn           net.Conn
	scanner        *bufio.Scanner
	messageHandler MessageHandler
}

func NewClient(endpoint string) (*client, error) {
	if endpoint == "" {
		log.Println("NewClient: missing endpoint")
		os.Exit(1)
	}
	conn, err := tls.Dial("tcp", endpoint, &tls.Config{})

	client := client{
		conn:    conn,
		scanner: bufio.NewScanner(conn),
	}

	return &client, err
}

func (client *client) ListenForMessages() {
	log.Println("Client.ListenForMessages()")
	for client.scanner.Scan() {
		line := client.scanner.Text()

		// handle messages which require a reply
		if strings.HasPrefix(line, "PING") {
			err := client.Send(strings.Replace(line, "PING", "PONG", 1))
			if err != nil {
				log.Println("Error sending PONG, we'll probably get terminated soon")
			}
		}

		if client.messageHandler != nil {
			// @todo refactor this to a Message
			client.messageHandler(line)
		} else {
			log.Println("no messageHandler set!")
			log.Println(line)
		}
	}
}

func (client *client) Send(message string) error {
	log.Print("Client.Send() -> ", message)
	_, err := client.conn.Write([]byte(message))
	return err
}

func (client *client) OnMessage(handler MessageHandler) {
	client.messageHandler = handler
}

func (client *client) Authenticate(username string, password string) error {
	log.Printf("Client.Authenticate(%s)\n", username)
	pass := fmt.Sprintf("PASS oauth:%s\n", password)
	nick := fmt.Sprintf("NICK %s\n", username)

	err := client.Send(pass)
	if err != nil {
		log.Println("ERROR: Authenticate (PASS): ", err.Error())
		return err
	}
	err = client.Send(nick)
	if err != nil {
		log.Println("ERROR: Authenticate (NICK):", err.Error())
		return err
	}
	return nil
}

func (client *client) JoinChannel(channel string) error {
	log.Printf("Client.JoinChannel(%s)\n", channel)
	nick := fmt.Sprintf("JOIN %s", channel)
	err := client.Send(nick)
	if err != nil {
		log.Println("ERROR: JoinChannel (JOIN):", err.Error())
		return err
	}
	return nil
}

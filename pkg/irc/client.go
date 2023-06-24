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

type ReplyCallback func(command string)

type MessageHandler func(message Message, replyWith ReplyCallback)

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
		// log.Printf("----------(%s)----------", line)

		if client.messageHandler == nil {
			log.Println("no messageHandler set!")
			log.Println(line)
			continue
		}

		message, err := NewMessage(line)
		if err != nil {
			log.Printf("failed to parse message (%s)", message.Text)
			continue
		}

		// log.Println("Client.Rcvd() <- ", message.Command.Channel)
		if message.Command.Command == "PING" {
			client.Pong(message.Parameters)
		} else {
			client.messageHandler(message, func(replyWith string) {
				if replyWith != "" {
					client.Send(replyWith)
				}
			})
		}
	}
}

func (client *client) Send(message string) error {
	if strings.Contains(message, "PASS") {
		log.Println("Client.Send() -> PASS:REDACTED")
	} else {
		log.Println("Client.Send() -> ", message)
	}
	if !strings.HasSuffix(message, "\r\n") {
		message = fmt.Sprintf("%s\r\n", message)
	}
	_, err := client.conn.Write([]byte(message))
	return err
}

func (client *client) OnMessage(handler MessageHandler) {
	client.messageHandler = handler
}

func (client *client) RegisterBotCommand() {
	// Future stuff?
}

func (client *client) Pong(parameters string) {
	cmd := fmt.Sprintf("PONG %s", parameters)
	err := client.Send(cmd)
	if err != nil {
		log.Println("ERROR: ", cmd, err.Error())
	}
}

func (client *client) Authenticate(username string, password string) error {
	log.Printf("Client.Authenticate(%s)\n", username)
	pass := fmt.Sprintf("PASS oauth:%s", password)
	nick := fmt.Sprintf("NICK %s", username)

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

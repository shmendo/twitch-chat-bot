package irc

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type ReplyCallback func(command string)

type MessageHandler func(message Message, replyWith ReplyCallback)

type client struct {
	conn       net.Conn
	scanner    *bufio.Scanner
	sendBuffer chan string

	// message specific handlers
	onMessage        MessageHandler
	onAuthenticated  MessageHandler
	onPrivateMessage MessageHandler
	onBotCommand     MessageHandler
}

func NewClient(endpoint string) (*client, error) {
	if endpoint == "" {
		log.Println("NewClient: missing endpoint")
		os.Exit(1)
	}
	conn, err := tls.Dial("tcp", endpoint, &tls.Config{})

	// var logBuffer bytes.Buffer
	// logger := log.New(&logBuffer, "client: ", log.Lshortfile)
	// logger.Print("Hello, log file!")
	// output: "logger: client.go:19: Hello, log file!"

	client := client{
		conn:       conn,
		scanner:    bufio.NewScanner(conn),
		sendBuffer: make(chan string, 100),
	}

	return &client, err
}

func (client *client) ListenForMessages() {
	for client.scanner.Scan() {
		line := client.scanner.Text()

		message, err := NewMessage(line)
		if err != nil {
			log.Printf("failed to parse message (%s)", message.Text)
			continue
		}
		switch message.Command.Command {
		case "PING":
			client.Pong(message.Parameters)
		case "376":
			if client.onAuthenticated != nil {
				client.onAuthenticated(message, client.queueMessage)
			}
		case "JOIN":
			log.Printf("Successfully joined %s", message.Command.Channel)
		case "PRIVMSG":
			if message.CommandType == "bot" && client.onBotCommand != nil {
				client.onBotCommand(message, client.queueMessage)
			} else {
				client.onPrivateMessage(message, client.queueMessage)
			}
		}
		if client.onMessage != nil {
			client.onMessage(message, client.queueMessage)
		}
	}
}

func (client *client) queueMessage(message string) {
	log.Printf("Client->queueMessage(%s)", message)
	if message != "" {
		client.sendBuffer <- message
	}

	// we can send up to 100 messages every 30 seconds
	// which equates to around 1 messages every 0.3s
	// we don't want to hit the limit though, so we'll
	// go just a little slower (.27s)
	go func() {
		for messageToSend := range client.sendBuffer {
			time.Tick(time.Millisecond * 300)
			err := client.Send(messageToSend)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

func (client *client) Send(message string) error {
	log.Printf("Client->Send(%s)", message)
	if !strings.HasSuffix(message, "\r\n") {
		message = fmt.Sprintf("%s\r\n", message)
	}
	_, err := client.conn.Write([]byte(message))
	return err
}

func (client *client) OnMessage(handler MessageHandler) {
	log.Println("Client->OnMessage")
	client.onMessage = handler
}

func (client *client) OnAuthenticated(handler MessageHandler) {
	log.Println("Client->OnAuthenticated")
	client.onAuthenticated = handler
}

func (client *client) OnPrivateMessage(handler MessageHandler) {
	log.Println("Client->OnPrivateMessage")
	client.onPrivateMessage = handler
}

func (client *client) OnBotCommand(handler MessageHandler) {
	log.Println("Client->OnBotCommand")
	client.onBotCommand = handler
}

func (client *client) Pong(parameters string) {
	cmd := fmt.Sprintf("Client->PONG %s", parameters)
	err := client.Send(cmd)
	if err != nil {
		log.Println("ERROR: ", cmd, err.Error())
	}
}

func (client *client) Authenticate(username string, password string) {
	log.Printf("Client->Authenticate(%s)\n", username)
	pass := fmt.Sprintf("PASS oauth:%s", password)
	nick := fmt.Sprintf("NICK %s", username)

	client.queueMessage(pass)
	client.queueMessage(nick)
}

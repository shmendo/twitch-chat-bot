package irc

import (
	"log"
	"strings"
)

// @see https://ircv3.net/specs/extensions/message-tags.html#format
type Tag struct {
	Key   string
	Value string
}

type Source struct {
	Nick string
	Host string
}

type Command struct {
	Command string
	Channel string
	Info    string
}

type Message struct {
	Text        string
	Valid       bool
	Tags        []Tag
	Source      Source
	Command     Command
	CommandType string
	Parameters  string
}

func NewMessage(messageText string) (Message, error) {
	// log.Println("NewMessage()", messageText)
	// initialize message w/original text
	message := Message{
		Text: messageText,
	}

	textToProcess := messageText
	idx := 0
	segment := ""

	if string(textToProcess[0]) == "@" {
		idx = strings.Index(textToProcess, " ")
		segment = textToProcess[0:idx]
		addTags(&message, segment)
		textToProcess = textToProcess[idx:]
	}
	// log.Println("NewMessage(Tags) > ", 0, idx, segment)

	// Extract SOURCE (nick@host)
	// The idx should point to the source part; otherwise, it's a PING command.
	if string(textToProcess[0]) == ":" {
		idx = strings.Index(textToProcess, " ")
		segment = textToProcess[1:idx]
		addSource(&message, segment)
		textToProcess = textToProcess[idx:]
	}
	// log.Println("NewMessage(Source) > ", 0, idx, segment)

	// Extract COMMAND
	idx = strings.Index(textToProcess, ":")
	if idx == -1 {
		// Not all messages include parameters
		idx = len(textToProcess)
	}
	segment = textToProcess[0:idx]
	// log.Println("NewMessage(Command) > ", 0, idx, segment)
	addCommand(&message, segment)
	textToProcess = textToProcess[idx:]

	// Extract PARAMETERS
	if len(textToProcess) > 0 {
		segment = textToProcess[1:]
		// log.Println("NewMessage(Parameters) > ", 0, idx, segment)
		addParameters(&message, segment)
	}

	return message, nil
}

// Parse the tagString into []Tag
func addTags(message *Message, tagString string) {
	// pretty sure we won't be using these, but basic stuff setup
	tagSegments := strings.Split(tagString, ";")
	for _, tagSegment := range tagSegments {
		t := strings.Split(tagSegment, "=")
		log.Printf("%s=%s", t[0], t[1])
		// append(message.Tags, Tag{
		// 	key:   t[0],
		// 	value: t[1],
		// })
	}
}

// Parse out source
func addSource(message *Message, sourceString string) {
	sourceSegments := strings.Split(sourceString, "!")
	message.Source.Nick = ""
	if len(sourceSegments) == 2 {
		message.Source.Nick = sourceSegments[0]
		message.Source.Host = sourceSegments[1]
	} else {
		message.Source.Host = sourceSegments[0]
	}
}

// Parse out command
func addCommand(message *Message, commandString string) {
	commandSegments := strings.Split(commandString, " ")
	message.Command.Command = commandSegments[0]
	message.Command.Channel = ""
	message.Command.Info = ""
	if len(commandSegments) >= 2 {
		message.Command.Channel = commandSegments[1]
	}
	if len(commandSegments) >= 3 {
		message.Command.Info = commandSegments[2]
	}
}

// Parse out parameters
func addParameters(message *Message, parameterString string) {
	message.CommandType = "standard"

	if string(parameterString[0]) == "!" {
		message.CommandType = "bot"
	}
}

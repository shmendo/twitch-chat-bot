package irc

import (
	"errors"
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
	log.Println("NewMessage()", messageText)
	// initialize message w/original text
	message := Message{
		Text: messageText,
	}

	index := 0
	endIndex := 0
	segment := ""

	if string(messageText[index]) == "@" {
		endIndex = strings.Index(messageText, " ")
		if endIndex <= index {
			message.Valid = false
			return message, errors.New("invalid tags component")
		}
		segment = messageText[index:endIndex]
		addTags(&message, segment)
	}
	log.Println("NewMessage(Tags) > ", index, endIndex, segment)

	// Extract SOURCE (nick@host)
	// The idx should point to the source part; otherwise, it's a PING command.
	if string(messageText[index]) == ":" {
		index += 1 // discard :
		endIndex = strings.Index(messageText, " ")
		if endIndex <= index {
			message.Valid = false
			return message, errors.New("invalid source component")
		}
		segment = messageText[index:endIndex]
		addSource(&message, segment)
		index = endIndex + 1 // advance to next segment
	}
	log.Println("NewMessage(Source) > ", index, endIndex, segment)

	// Extract COMMAND
	endIndex = strings.Index(messageText, ":")
	if endIndex == -1 {
		// Not all messages include parameters
		endIndex = len(messageText)
	}
	segment = messageText[index:endIndex]
	log.Println("NewMessage(Command) > ", index, endIndex, segment)
	addCommand(&message, segment)

	// Extract PARAMETERS
	if endIndex != len(messageText) {
		index += 1
		segment = messageText[index:endIndex]
		log.Println("NewMessage(Parameters) > ", index, endIndex, segment)
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

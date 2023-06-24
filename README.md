# twitch-chat-bot

## Functionality

- define a command (message) to listen for from twitch chat
- provide authentication mechanism w/twitch so we can read messages on the chat window
- reach out to 3rd party API and then post some data back to the calling user
- should appear in the chat window

# Running

- `cp .env.example .env`
- update `.env` with appropriate values
- execute `source .env && go run cmd/chat-bot/main.go`

# Testing

I dropped a few basic tests to ensure message parsing was working as expected. Would liked to have gotten further :(

```
$ go test pkg/irc/*.go
```

# Notes/Investigation

## Define a command to listen for

```
!norris
```

## Authentication mechanism for reading/posting messages to twitch chat room

- IRC Protocol
- endpoint: irc://irc.chat.twitch.tv:6697 <-- SSL
- endpoint: irc://irc.chat.twitch.tv:6667 <-- NON-SSL


### IRC Messages to be implemented

- IRC RESERVED WORDS
- delimeted by \r\n
- PING - must respond with PONG or server will terminate connection
- 

### IRC Connection Flow

- connect to server
- PASS oauth:{API_TOKEN} (see https://dev.twitch.tv/docs/irc/authenticate-bot/)
- JOIN #channel_name - responds with 353 & 366 messages (see https://dev.twitch.tv/docs/irc/join-chat-room/)

### Receiving chat messages
- :foo!foo@foo.tmi.twitch.tv PRIVMSG #bar :bleedPurple

### Sending chat messages
- PRIVMSG #channel_name :This is a sample message
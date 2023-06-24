# twitch-chat-bot

## Functionality

- define a command (message) to listen for from twitch chat
- provide authentication mechanism w/twitch so we can read messages ion the chat window
- reach out to 3rd party API and then post some data back
- post response to the chat window


## Define a command to listen for

```
!busey
```

## Authentication mechanism for reading/posting messages to twitch chat room

- IRC Protocol
- endpoint: irc://irc.chat.twitch.tv:6697


### IRC Messages to be implemented

- IRC RESERVED WORDS
- delimeted by \r\n
- 
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
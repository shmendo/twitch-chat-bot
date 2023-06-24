package irc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    Message
		wantError   bool
	}{
		{
			description: "should handle PING",
			input:       "PING :tmi.twitch.tv",
			expected: Message{
				Text:  "PING :tmi.twitch.tv",
				Valid: true,
				Tags:  []Tag(nil),
				Source: Source{
					Nick: "",
					Host: "",
				},
				Command: Command{
					Command: "PING",
					Channel: "",
					Info:    "",
				},
				CommandType: "standard",
				Parameters:  "tmi.twitch.tv",
			},
			wantError: false,
		},
		// {
		// 	description: "should handle !norris",
		// 	input:       "PING :tmi.twitch.tv",
		// 	expected: Message{
		// 		Text:  "PING :tmi.twitch.tv",
		// 		Valid: true,
		// 		Tags:  []Tag(nil),
		// 		Source: Source{
		// 			Nick: "",
		// 			Host: "",
		// 		},
		// 		Command: Command{
		// 			Command: "PING",
		// 			Channel: "",
		// 			Info:    "",
		// 		},
		// 		CommandType: "standard",
		// 		Parameters:  "tmi.twitch.tv",
		// 	},
		// 	wantError: false,
		// }
		// {
		// 	description: "should p",
		// 	input:       ":tmi.twitch.tv 001 benjamin_walters :Welcome, GLHF!",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":tmi.twitch.tv 002 benjamin_walters :Your host is tmi.twitch.tv",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":tmi.twitch.tv 003 benjamin_walters :This server is rather new",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":tmi.twitch.tv 004 benjamin_walters :-",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":tmi.twitch.tv 375 benjamin_walters :-",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":tmi.twitch.tv 372 benjamin_walters :You are in a maze of twisty passages, all alike.",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":tmi.twitch.tv 376 benjamin_walters :>",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":benjamin_walters!benjamin_walters@benjamin_walters.tmi.twitch.tv JOIN #benjamin_walters",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":benjamin_walters.tmi.twitch.tv 353 benjamin_walters = #benjamin_walters :benjamin_walters",
		// 	expected:    Message{},
		// 	wantError:   false,
		// }, {
		// 	description: "",
		// 	input:       ":benjamin_walters.tmi.twitch.tv 366 benjamin_walters #benjamin_walters :End of /NAMES list",
		// 	expected:    Message{},
		// 	wantError:   false,
		// },
	}

	for _, c := range cases {
		t.Run(
			c.description,
			func(t *testing.T) {
				actual, err := NewMessage(c.input)
				if c.wantError {
					assert.Error(t, err, fmt.Sprintf("unexpected value returned %v", actual))
					return
				}
				assert.NoError(t, err, actual)
				assert.Equal(t, c.expected, actual)
			},
		)
	}
}

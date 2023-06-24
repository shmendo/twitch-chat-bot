package irc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMessage(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    Message
		wantError   bool
	}{
		{
			"should handle an empty message",
			"",
			Message{},
			true,
		},
		{
			"should handle an invalid message",
			"neeners my bruv",
			Message{},
			true,
		},
		{
			"should parse welcome message",
			":tmi.twitch.tv 001 benjamin_walters :Welcome, GLHF!",
			Message{},
			false,
		},
		{
			"should parse ",
			":benjamin_walters.tmi.twitch.tv 366 benjamin_walters #benjamin_walters :End of /NAMES list",
			Message{},
			false,
		},
		{
			"should handle special PING command",
			"PING :tmi.twitch.tv",
			Message{},
			false,
		},
	}

	for _, c := range cases {
		t.Run(
			c.description,
			func(t *testing.T) {
				assert.NoError(t)
			},
		)
	}
}

package irc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
		wantError   bool
	}{}

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

func TestListenForMessages(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
		wantError   bool
	}{}

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

func TestSend(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
		wantError   bool
	}{}

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

func TestOnMessage(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
		wantError   bool
	}{}

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

func TestRegisterBotCommand(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
		wantError   bool
	}{}

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

func TestPong(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
		wantError   bool
	}{}

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

func TestAuthenticate(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expected    string
		wantError   bool
	}{}

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

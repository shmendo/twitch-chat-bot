package irc

type Message struct {
	prefix     string
	command    string
	parameters []string
}

// @see https://datatracker.ietf.org/doc/html/rfc1459.html#section-6
type Reply struct {
}

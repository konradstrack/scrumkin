package ping

import (
	"fmt"
	"log"
	"regexp"
	"scrumkin/pkg/messages"
)

type Ping struct {
	Name        string
	MatchRegexp *regexp.Regexp
}

func New() (*Ping, error) {
	name := "ping"
	r, err := regexp.Compile(`^ping\s*$`)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, fmt.Errorf("Cannot create command: %s", name)
	}

	return &Ping{name, r}, nil
}

func (cmd *Ping) Match(msg *messages.Message) bool {
	return cmd.MatchRegexp.MatchString(msg.Text)
}

func (cmd *Ping) Process(msg *messages.Message) messages.Response {
	return messages.Response{"pong"}
}

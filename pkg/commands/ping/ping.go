package ping

import (
	"fmt"
	"log"
	"regexp"
	"scrumkin/pkg/messages"
)

type Ping struct {
	name        string
	matchRegexp *regexp.Regexp
	helpSummary string
}

func New() (*Ping, error) {
	name := "ping"
	r, err := regexp.Compile(`(?i)^\s*ping\s*$`)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, fmt.Errorf("Cannot create command: %s", name)
	}

	summary := "Returns 'pong'"
	return &Ping{name, r, summary}, nil
}

func (p *Ping) Match(msg *messages.Message) bool {
	return p.matchRegexp.MatchString(msg.Text)
}

func (p *Ping) Process(msg *messages.Message) *messages.Response {
	return &messages.Response{"pong"}
}

func (p *Ping) Name() string {
	return p.name
}

func (p *Ping) HelpSummary() string {
	return p.helpSummary
}

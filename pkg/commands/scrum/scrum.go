package scrum

import (
	"fmt"
	"log"
	"regexp"
	"scrumkin/pkg/messages"
	"scrumkin/pkg/nlp"
)

// Scrum represents the scrum command
type Scrum struct {
	name        string
	matchRegexp *regexp.Regexp
	helpSummary string
}

// New returns a new scrum command
func New() (*Scrum, error) {
	name := "scrum"
	r, err := regexp.Compile(`(?i)^\s*(scrum|today|yesterday)\s+.*$`)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, fmt.Errorf("Cannot create command: %s", name)
	}

	summary := "Lets you scrum"
	return &Scrum{name, r, summary}, nil
}

// Match determines whether a given message matches the command
func (p *Scrum) Match(msg *messages.Message) bool {
	return p.matchRegexp.MatchString(msg.Text)
}

// Process handles the command and creates a response
func (p *Scrum) Process(msg *messages.Message) *messages.Response {
	config := &nlp.Config{
		BaseURL:    "http://corenlp:9000",
		Annotators: []string{"tokenize", "ssplit", "pos", "depparse", "lemma", "ner"},
	}
	client := nlp.NewCoreNLPClient(config)
	log.Print(client.Query(msg.Text))
	return &messages.Response{"Not implemented yet"}
}

// Name returns command's name
func (p *Scrum) Name() string {
	return p.name
}

// HelpSummary returns short help summary for the command
func (p *Scrum) HelpSummary() string {
	return p.helpSummary
}

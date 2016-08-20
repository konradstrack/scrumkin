package scrum

import (
	"fmt"
	"regexp"
	"scrumkin/pkg/messages"
	"scrumkin/pkg/nlp"
)

// Scrum represents the scrum command
type Scrum struct {
	name        string
	matchRegexp *regexp.Regexp
	helpSummary string
	nlpClient   nlp.Client
}

// New returns a new scrum command
func New() (*Scrum, error) {
	name := "scrum"
	r, err := regexp.Compile(`(?i)^\s*(scrum|today|yesterday).*$`)
	if err != nil {
		return nil, fmt.Errorf("Cannot create command %s: %s", name, err)
	}

	config := &nlp.CoreNLPConfig{
		BaseURL:    "http://corenlp:9000",
		Annotators: []string{"tokenize", "ssplit", "pos", "depparse", "lemma", "ner"},
	}
	client := nlp.NewCoreNLPClient(config)

	summary := "Lets you scrum"
	return &Scrum{name, r, summary, client}, nil
}

// Match determines whether a given message matches the command
func (s *Scrum) Match(msg *messages.Message) bool {
	return s.matchRegexp.MatchString(msg.Text)
}

// Process handles the command and creates a response
func (s *Scrum) Process(msg *messages.Message) *messages.Response {
	return &messages.Response{"Not implemented yet"}
}

// Name returns command's name
func (s *Scrum) Name() string {
	return s.name
}

// HelpSummary returns short help summary for the command
func (s *Scrum) HelpSummary() string {
	return s.helpSummary
}

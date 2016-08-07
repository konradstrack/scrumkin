package help

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"scrumkin/pkg/commands"
	"scrumkin/pkg/messages"
)

type Help struct {
	name        string
	matchRegexp *regexp.Regexp
	helpSummary string
	commands    []commands.Command
}

func New() (*Help, error) {
	name := "help"
	r, err := regexp.Compile(`(?i)^\s*help\s*$`)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, fmt.Errorf("Cannot create command: %s", name)
	}

	summary := "Displays help"
	return &Help{name, r, summary, nil}, nil
}

func (h *Help) SetCommands(cmds []commands.Command) {
	h.commands = cmds
}

func (h *Help) Match(msg *messages.Message) bool {
	return h.matchRegexp.MatchString(msg.Text)
}

func (h *Help) Process(msg *messages.Message) *messages.Response {
	var buffer bytes.Buffer

	if h.commands != nil {
		fmt.Fprintf(&buffer, "Available commands:\n")
		for _, cmd := range h.commands {
			fmt.Fprintf(&buffer, "\t*%s*: %s\n", cmd.Name(), cmd.HelpSummary())
		}
	} else {
		fmt.Fprintf(&buffer, "No help available")
		log.Printf("error: no commands registerd for help")
	}

	log.Printf(buffer.String())
	return &messages.Response{buffer.String()}
}

func (h *Help) Name() string {
	return h.name
}

func (h *Help) HelpSummary() string {
	return h.helpSummary
}

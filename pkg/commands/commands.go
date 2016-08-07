package commands

import "scrumkin/pkg/messages"

type Command interface {
	Match(*messages.Message) bool
	Process(*messages.Message) *messages.Response
	Name() string
	HelpSummary() string
}

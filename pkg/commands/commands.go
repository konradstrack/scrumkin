package commands

import "scrumway/pkg/messages"

type Command interface {
	Match(*messages.Message) bool
	Process(*messages.Message) messages.Response
}

package bot

import (
	"log"
	"scrumkin/pkg/commands/help"
	"scrumkin/pkg/commands/ping"
	"scrumkin/pkg/messages"
)

func (b *Bot) registerCommands() {
	// register ping
	pingCmd, err := ping.New()
	if err != nil {
		log.Fatalf("fatal error while registering '%s': %s", pingCmd.Name(), err)
	}

	b.commands = append(b.commands, pingCmd)

	// register help
	helpCmd, err := help.New()
	if err != nil {
		log.Fatalf("fatal error while registering 'help' command: %s", err)
	}

	b.commands = append(b.commands, helpCmd)
	helpCmd.SetCommands(b.commands)
}

func (b *Bot) handleMessages() {
	for {
		msg := <-b.messages
		log.Print(msg)
		b.processMessage(&msg)
	}
}

func (b *Bot) processMessage(msg *messages.Message) {
	for _, cmd := range b.commands {
		if cmd.Match(msg) {
			response := cmd.Process(msg)
			b.sendResponse(response, msg)
		}
	}
}

func (b *Bot) sendResponse(r *messages.Response, oldMsg *messages.Message) {
	msg := b.rtm.NewOutgoingMessage(r.Text, oldMsg.Channel)
	b.rtm.SendMessage(msg)
}

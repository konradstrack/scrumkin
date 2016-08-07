package bot

import (
	"log"
	"scrumway/pkg/commands/ping"
	"scrumway/pkg/messages"
)

func (b *Bot) registerCommands() {
	pingCmd, err := ping.New()
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	b.commands = append(b.commands, pingCmd)
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
			log.Print(cmd.Process(msg))
		}
	}
}

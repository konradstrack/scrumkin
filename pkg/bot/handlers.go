package bot

import (
	"log"
	"scrumkin/pkg/commands"
	"scrumkin/pkg/commands/help"
	"scrumkin/pkg/commands/ping"
	"scrumkin/pkg/commands/scrum"
	"scrumkin/pkg/messages"
)

func (b *Bot) registerCommands() {
	pingCmd, pingErr := ping.New()
	scrumCmd, scrumErr := scrum.New()
	helpCmd, helpErr := help.New()

	cmds := []struct {
		cmd commands.Command
		err error
	}{
		{pingCmd, pingErr},
		{scrumCmd, scrumErr},
		{helpCmd, helpErr},
	}

	for _, c := range cmds {
		log.Printf("Registering command: %s", c.cmd.Name())
		if c.err != nil {
			log.Fatalf("fatal error: registering '%s': %s", c.cmd.Name(), c.err)
		}
		b.commands = append(b.commands, c.cmd)
	}

	// set commands for help
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

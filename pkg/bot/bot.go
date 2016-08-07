package bot

import (
	"log"
	"os"
	"scrumkin/pkg/commands"
	"scrumkin/pkg/messages"

	"github.com/nlopes/slack"
)

type Bot struct {
	Token    string
	messages chan messages.Message
	commands []commands.Command
	rtm      *slack.RTM
}

func (b *Bot) Run() {
	setUpLogger()
	b.registerCommands()

	go b.listen()
	b.handleMessages()
}

func New(token string) *Bot {
	api := slack.New(token)

	// Enable Slack api debugging if env variable set
	if os.Getenv("DEBUG") != "" {
		api.SetDebug(true)
	}

	rtm := api.NewRTM()

	bot := &Bot{
		Token:    token,
		messages: make(chan messages.Message),
		commands: make([]commands.Command, 0),
		rtm:      rtm,
	}

	return bot
}

func (b *Bot) Enqueue(m messages.Message) {
	b.messages <- m
}

func (b *Bot) listen() *slack.RTM {
	go b.rtm.ManageConnection()

	for {
		msg := <-b.rtm.IncomingEvents
		switch event := msg.Data.(type) {
		case *slack.ConnectedEvent:
			printConnectionInfo(event)
		case *slack.MessageEvent:
			m := messages.Message{
				Text:    event.Text,
				User:    event.User,
				Channel: event.Channel,
			}
			b.Enqueue(m)
		}
	}
}

func setUpLogger() {
	// added a logger to fix a nil pointer dereference;
	// see: https://github.com/nlopes/slack/commit/faac376828565b0d1dce05142add386de5fb7363
	logger := log.New(os.Stdout, "scrumkin: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
}

func printConnectionInfo(event *slack.ConnectedEvent) {
	name := event.Info.Team.Name
	domain := event.Info.Team.Domain
	log.Printf("Connected to team %s (%s)", name, domain)
}

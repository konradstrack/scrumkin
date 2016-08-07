package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

type Bot struct {
	Token    string
	Messages chan Message
}

func (b *Bot) Run() {
	setUpLogger()

	go b.listen()
	b.handleMessages()
}

func New(token string) *Bot {
	bot := &Bot{
		Token:    token,
		Messages: make(chan Message),
	}

	return bot
}

func (b *Bot) Enqueue(m Message) {
	b.Messages <- m
}

func (b *Bot) listen() *slack.RTM {
	api := slack.New(b.Token)

	// Enable Slack api debugging if env variable set
	if os.Getenv("DEBUG") != "" {
		api.SetDebug(true)
	}

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		msg := <-rtm.IncomingEvents
		switch event := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Connected:")
		case *slack.MessageEvent:
			m := Message{event.Msg.Text}
			b.Enqueue(m)
		}
	}
}

func setUpLogger() {
	// added a logger to fix a nil pointer dereference;
	// see: https://github.com/nlopes/slack/commit/faac376828565b0d1dce05142add386de5fb7363
	logger := log.New(os.Stdout, "scrumway: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
}

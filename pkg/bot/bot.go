package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

type Bot struct {
	messages chan *Message
}

type Message struct {
	message string
}

func Run() {
	setUpLogger()

	bot := new(Bot)
	bot.Connect()
}

func (b *Bot) Enqueue(m *Message) {
	b.messages <- m
}

func (b *Bot) Connect() *slack.RTM {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("error: Token not provided. Please set the SLACK_TOKEN environment variable.")
	}

	api := slack.New(token)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		msg := <-rtm.IncomingEvents
		switch event := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Connected:")
		case *slack.MessageEvent:
			m := Message{event.Msg.Text}
			b.Enqueue(&m)
		}
	}
}

func setUpLogger() {
	// added a logger to fix a nil pointer dereference;
	// see: https://github.com/nlopes/slack/commit/faac376828565b0d1dce05142add386de5fb7363
	logger := log.New(os.Stdout, "scrumway: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
}

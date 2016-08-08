package bot

import (
	"log"
	"os"
	"scrumkin/pkg/commands"
	"scrumkin/pkg/db"
	"scrumkin/pkg/messages"

	"github.com/nlopes/slack"
)

// Bot represents bot's internal data
type Bot struct {
	Token    string
	dbDsn    string
	messages chan messages.Message
	commands []commands.Command
	rtm      *slack.RTM
	db       *db.DB
}

// Run starts the bot
func (b *Bot) Run() {
	setUpLogger()

	err := b.connectToDatabase()
	if err != nil {
		log.Fatalf("fatal error: cannot connect to database: %s", err)
	}

	b.registerCommands()

	go b.listen()
	b.handleMessages()
}

// New creates a new bot and initializes default values
func New(token string, dbDsn string) *Bot {
	api := slack.New(token)

	// Enable Slack api debugging if env variable set
	if os.Getenv("DEBUG") != "" {
		api.SetDebug(true)
	}

	rtm := api.NewRTM()

	bot := &Bot{
		Token:    token,
		dbDsn:    dbDsn,
		messages: make(chan messages.Message),
		commands: make([]commands.Command, 0),
		rtm:      rtm,
	}

	return bot
}

// Enqueue adds a message to the processing queue
func (b *Bot) Enqueue(m messages.Message) {
	b.messages <- m
}

// listen starts listening to incoming messages
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

func (b *Bot) connectToDatabase() error {
	db, err := db.New(b.dbDsn)
	if err != nil {
		return err
	}

	b.db = db
	return nil
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

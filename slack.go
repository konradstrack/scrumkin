package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func Connect() *slack.RTM {
	setupLogger()

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		fmt.Println("Token not provided. Please set the SLACK_TOKEN environment variable.")
		os.Exit(1)
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
			fmt.Printf("Message: %v\n", event)
		default:
			fmt.Println("Other event...")
		}
	}
}

func setupLogger() {
	// added a logger to fix a nil pointer dereference;
	// see: https://github.com/nlopes/slack/commit/faac376828565b0d1dce05142add386de5fb7363
	logger := log.New(os.Stdout, "scrumway: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
}

package main

import (
	"fmt"
	"log"
	"os"
	"scrumway/pkg/bot"
)

func main() {
	fmt.Println("Starting scrumway")

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("error: Token not provided. Please set the SLACK_TOKEN environment variable.")
	}

	b := bot.New(token)
	b.Run()
}

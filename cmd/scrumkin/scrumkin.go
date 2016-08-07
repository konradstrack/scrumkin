package main

import (
	"fmt"
	"log"
	"os"
	"scrumkin/pkg/bot"
)

func main() {
	fmt.Println("Starting scrumkin")

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("error: Token not provided. Please set the SLACK_TOKEN environment variable.")
	}

	b := bot.New(token)
	b.Run()
}

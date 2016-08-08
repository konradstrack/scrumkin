package main

import (
	"log"
	"os"
	"scrumkin/pkg/bot"
)

func main() {
	log.Printf("Starting scrumkin")

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("fatal error: Token not provided. Please set the SLACK_TOKEN environment variable.")
	}

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("fatal error: Data source name not provided. Please set the DATABASE_DSN environment variable.")
	}

	b := bot.New(token, dsn)
	b.Run()
}

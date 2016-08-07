package bot

import "log"

type Message struct {
	message string
}

func (b *Bot) handleMessages() {
	for {
		msg := <-b.Messages
		log.Print(msg)
	}
}

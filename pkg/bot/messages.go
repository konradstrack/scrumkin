package bot

import "log"

type Message struct {
	Text    string
	User    string
	Channel string
}

func (b *Bot) handleMessages() {
	for {
		msg := <-b.Messages
		log.Print(msg)
	}
}

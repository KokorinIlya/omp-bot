package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (commander *ItemCommander) SendMessage(uid int64, messageText string) {
	msg := tgbotapi.NewMessage(uid, messageText)
	_, err := commander.botApi.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat: %v", err)
	}
}

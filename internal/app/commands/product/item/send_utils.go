package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (commander *ItemCommander) editMessage(msg tgbotapi.EditMessageTextConfig) {
	_, err := commander.botApi.Send(msg)
	if err != nil {
		log.Printf("Error editing text of message from chat %v: %v", msg.ChatID, err)
	}
}

func (commander *ItemCommander) sendMessage(msg tgbotapi.MessageConfig) {
	_, err := commander.botApi.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat %v: %v", msg.ChatID, err)
	}
}

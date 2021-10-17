package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ItemCommander) Default(inputMessage *tgbotapi.Message) {
	msgText := fmt.Sprintf("Unknown command %s, type /help__product__item for help", inputMessage.Text)
	commander.SendMessage(inputMessage.Chat.ID, msgText)
}

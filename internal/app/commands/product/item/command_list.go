package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ItemCommander) List(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	msgText, keyboard, err := getPaginatedMessage(commander.itemService, 0, DefaultItemsPerPage)
	if err != nil {
		errText := fmt.Sprintf("Cannot process /list command: %v", err)
		errMsg := tgbotapi.NewMessage(chatId, errText)
		commander.sendMessage(errMsg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, msgText)
	msg.ReplyMarkup = keyboard
	commander.sendMessage(msg)
}

package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ItemCommander) Get(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()
	itemId, err := commander.parseIdOrSendError(commandArgs, chatId, "to get item by")
	if err != nil {
		return
	}
	item, err := commander.itemService.Describe(itemId)
	if err != nil {
		msgText := fmt.Sprintf("Cannot get item by id %v: %v", itemId, err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, item.String())
	commander.sendMessage(msg)
}

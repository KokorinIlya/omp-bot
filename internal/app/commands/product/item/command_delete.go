package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ItemCommander) Delete(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()
	itemId, err := commander.parseIdOrSendError(commandArgs, chatId, "to delete item by")
	if err != nil {
		return
	}
	err = commander.itemService.Remove(itemId)
	if err != nil {
		msgText := fmt.Sprintf("Cannot delete item by id %v: %v", itemId, err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, "Successfully removed item")
	commander.sendMessage(msg)
}

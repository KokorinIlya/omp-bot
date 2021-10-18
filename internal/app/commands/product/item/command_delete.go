package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (commander *ItemCommander) Delete(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()
	itemIndex, err := strconv.Atoi(commandArgs)
	if err != nil {
		msgText := fmt.Sprintf("Couldn't parse index to delete item by: %v", err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	if itemIndex < 0 {
		msgText := fmt.Sprintf("Expected item index not less than zero, but %v received", itemIndex)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	err = commander.itemService.Remove(uint64(itemIndex))
	if err != nil {
		msgText := fmt.Sprintf("Cannot delete item by index %v: %v", itemIndex, err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, "Successfully removed item")
	commander.sendMessage(msg)
}

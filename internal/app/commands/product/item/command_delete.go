package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (commander *ItemCommander) Delete(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()
	itemId, err := strconv.ParseUint(commandArgs, 10, 64)
	if err != nil {
		msgText := fmt.Sprintf("Couldn't parse id to delete item by: %v", err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
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

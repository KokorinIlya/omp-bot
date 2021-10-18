package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (commander *ItemCommander) Get(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()
	itemIndex, err := strconv.Atoi(commandArgs)
	if err != nil {
		msgText := fmt.Sprintf("Couldn't parse index to get item by: %v", err)
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
	item, err := commander.itemService.Describe(uint64(itemIndex))
	if err != nil {
		msgText := fmt.Sprintf("Cannot get item by index %v: %v", itemIndex, err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, item.String())
	commander.sendMessage(msg)
}

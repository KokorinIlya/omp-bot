package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
	"strconv"
	"strings"
)

func (commander *ItemCommander) Edit(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := strings.Split(inputMessage.CommandArguments(), " ")
	if len(commandArgs) != 2 {
		msgText := fmt.Sprintf(
			"Expected 2 arguments: <index> <title>, but received %v: %v",
			len(commandArgs), commandArgs,
		)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	itemIndex, err := strconv.Atoi(commandArgs[0])
	if err != nil {
		msgText := fmt.Sprintf("Couldn't parse index to edit item by: %v", err)
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
	newTitle := commandArgs[1]
	newItem := item.NewItem(newTitle)
	err = commander.itemService.Update(uint64(itemIndex), *newItem)
	if err != nil {
		msgText := fmt.Sprintf("Cannot edit item by index %v: %v", itemIndex, err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, "Successfully updated item")
	commander.sendMessage(msg)
}

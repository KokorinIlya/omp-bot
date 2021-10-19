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
			"Expected 2 arguments: <id> <title>, but received %v: %v",
			len(commandArgs), commandArgs,
		)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	itemId, err := strconv.ParseUint(commandArgs[0], 10, 64)
	if err != nil {
		msgText := fmt.Sprintf("Couldn't parse id to edit item by: %v", err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	newTitle := commandArgs[1]
	newItem := item.NewItem(itemId, newTitle)
	err = commander.itemService.Update(itemId, *newItem)
	if err != nil {
		msgText := fmt.Sprintf("Cannot edit item by id %v: %v", itemId, err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, "Successfully updated item")
	commander.sendMessage(msg)
}

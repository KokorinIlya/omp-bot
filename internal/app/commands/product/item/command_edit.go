package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
	"strings"
)

func (commander *ItemCommander) Edit(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := strings.Split(inputMessage.CommandArguments(), " ")
	if !commander.validateArgumentsCountOrSendError(commandArgs, 4, chatId,
		"<id> <owner_id> <product_id> <title>") {
		return
	}
	itemId, err := commander.parseIdOrSendError(commandArgs[0], chatId, "to edit item by")
	if err != nil {
		return
	}
	ownerId, err := commander.parseIdOrSendError(commandArgs[1], chatId, "of owner")
	if err != nil {
		return
	}
	productId, err := commander.parseIdOrSendError(commandArgs[2], chatId, "of product")
	if err != nil {
		return
	}

	newTitle := commandArgs[3]
	newItem := item.NewItem(itemId, ownerId, productId, newTitle)

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

package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
	"strings"
)

func (commander *ItemCommander) New(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	commandArgs := strings.Split(inputMessage.CommandArguments(), " ")
	if !commander.validateArgumentsCountOrSendError(commandArgs, 3, chatId,
		"<owner_id> <product_id> <title>") {
		return
	}
	ownerId, err := commander.parseIdOrSendError(commandArgs[0], chatId, "of owner")
	if err != nil {
		return
	}
	productId, err := commander.parseIdOrSendError(commandArgs[1], chatId, "of product")
	if err != nil {
		return
	}
	title := commandArgs[2]


	// Ids are allocated by Service
	newItem := item.NewItem(0, ownerId, productId, title)
	newId, err := commander.itemService.Create(*newItem)
	if err != nil {
		msgText := fmt.Sprintf("Couldn't create new item: %v", err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return
	}
	msgText := fmt.Sprintf("Successfully inserted new item, new item id is %v", newId)
	msg := tgbotapi.NewMessage(chatId, msgText)
	commander.sendMessage(msg)
}

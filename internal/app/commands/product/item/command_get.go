package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (commander *ItemCommander) Get(inputMessage *tgbotapi.Message) {
	uid := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()
	itemIndex, err := strconv.Atoi(commandArgs)
	if err != nil {
		msgText := fmt.Sprintf("Couldn't parse index to get item by: %v", err)
		commander.SendMessage(uid, msgText)
		return
	}
	item, err := commander.itemService.Describe(uint64(itemIndex))
	if err != nil {
		msgText := fmt.Sprintf("Cannot get item by index %v: %v", itemIndex, err)
		commander.SendMessage(uid, msgText)
		return
	}
	commander.SendMessage(uid, item.String())
}

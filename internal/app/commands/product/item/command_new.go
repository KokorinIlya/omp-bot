package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
)

func (commander *ItemCommander) New(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	title := inputMessage.CommandArguments()
	if title == "" {
		msg := tgbotapi.NewMessage(chatId, "Expected new item title")
		commander.sendMessage(msg)
		return
	}
	// Ids are allocated by ItemService
	newItem := item.NewItem(0, title)
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

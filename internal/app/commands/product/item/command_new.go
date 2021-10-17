package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
)

func (commander *ItemCommander) New(inputMessage *tgbotapi.Message) {
	uid := inputMessage.Chat.ID
	title := inputMessage.CommandArguments()
	if title == "" {
		commander.SendMessage(uid, "Expected new item title")
		return
	}
	newItem := item.NewItem(title)
	newIdx, err := commander.itemService.Create(*newItem)
	if err != nil {
		msgText := fmt.Sprintf("Coudn't create new item: %v", err)
		commander.SendMessage(uid, msgText)
		return
	}
	msgText := fmt.Sprintf("Successfully inserted new item, index is %v", newIdx)
	commander.SendMessage(uid, msgText)
}

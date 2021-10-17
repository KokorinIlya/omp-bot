package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ItemCommander) Help(inputMessage *tgbotapi.Message) {
	helpMsg := "/help__product__item - get help\n" +
		"/new__product__item <title> - create new item and receive newly-created item's index\n" +
		"/get__product__item <index> - get item by index"
	commander.SendMessage(inputMessage.Chat.ID, helpMsg)
}

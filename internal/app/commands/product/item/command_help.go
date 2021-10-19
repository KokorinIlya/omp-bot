package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ItemCommander) Help(inputMessage *tgbotapi.Message) {
	helpMsg := "/help__product__item - get help\n" +
		"/new__product__item <owner_id> <product_id> <title> - create new item and receive newly-created item's id\n" +
		"/get__product__item <id> - get item by id\n" +
		"/edit__product__item <id> <owner_id> <product_id> <title> - edit item by id\n" +
		"/delete__product__item <id> - delete item by id\n" +
		"/list__product__item - list all items with pagination"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpMsg)
	commander.sendMessage(msg)
}

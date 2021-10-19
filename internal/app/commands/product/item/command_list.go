package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (commander *ItemCommander) List(inputMessage *tgbotapi.Message) {
	chatId := inputMessage.Chat.ID
	msgText, keyboard, err := getPaginatedMessage(commander.itemService, 0, DefaultItemsPerPage)
	if err != nil {
		msgText := fmt.Sprintf("Cannot process /list command: %v", err)
		errMsg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(errMsg)
		return
	}
	msg := tgbotapi.NewMessage(chatId, msgText)
	msg.ReplyMarkup = keyboard
	_, err = commander.botApi.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat with id %v: %v", chatId, err)
	}
}

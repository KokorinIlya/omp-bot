package item

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (commander *ItemCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	chatId := callback.Message.Chat.ID
	var cursorData CursorData
	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &cursorData); err != nil {
		// TODO: maybe, send error to user
		log.Printf("Cannot parse cursor data from user message: %v", err)
		return
	}

	msgText, keyboard, err := getPaginatedMessage(commander.itemService,
		cursorData.Cursor, DefaultItemsPerPage)
	if err != nil {
		log.Printf("Error when getting paginated data: %v", err)
		return
	}
	msg := tgbotapi.NewEditMessageText(
		chatId,
		callback.Message.MessageID,
		msgText,
	)
	msg.ReplyMarkup = keyboard
	commander.editMessage(msg)
}

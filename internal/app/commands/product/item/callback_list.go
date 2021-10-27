package item

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (commander *ItemCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	chatId := callback.Message.Chat.ID
	var cursorData CursorData
	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &cursorData); err != nil {
		errText := fmt.Sprintf("Cannot parse cursor data from user message: %v", err)
		msg := tgbotapi.NewMessage(chatId, errText)
		commander.sendMessage(msg)
		return
	}

	msgText, keyboard, err := getPaginatedMessage(commander.itemService,
		cursorData.Cursor, DefaultItemsPerPage)
	if err != nil {
		errText := fmt.Sprintf("Error when getting paginated data: %v", err)
		msg := tgbotapi.NewMessage(chatId, errText)
		commander.sendMessage(msg)
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

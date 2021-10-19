package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (commander *ItemCommander) parseIdOrSendError(strId string, chatId int64, idDescription string) (uint64, error) {
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		msgText := fmt.Sprintf("Couldn't parse id %v: %v", idDescription, err)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return 0, err
	}
	return id, nil
}

func (commander *ItemCommander) validateArgumentsCountOrSendError(arguments []string, expectedLen int,
	chatId int64, argsDescription string,
) bool {
	if len(arguments) != expectedLen {
		msgText := fmt.Sprintf(
			"Expected %v arguments: %v, but received %v: %v",
			expectedLen, argsDescription, len(arguments), arguments,
		)
		msg := tgbotapi.NewMessage(chatId, msgText)
		commander.sendMessage(msg)
		return false
	}
	return true
}

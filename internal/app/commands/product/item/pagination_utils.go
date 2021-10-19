package item

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
	"strings"
)

const DefaultItemsPerPage = 3

type CursorData struct {
	Cursor uint64 `json:"offset"`
}

func getPaginatedMessage(itemService ItemService,
	cursor uint64, limit uint64) (string, *tgbotapi.InlineKeyboardMarkup, error) {
	items, err := itemService.List(cursor, limit)
	if err != nil {
		return "", nil, err
	}
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	if cursor > 0 {
		var newCursor uint64
		if limit > cursor {
			newCursor = 0
		} else {
			newCursor = cursor - limit
		}
		button, err := makeButton("К предыдущей странице", newCursor)
		if err != nil {
			return "", nil, err
		}
		buttons = append(buttons, *button)
	}
	if cursor+limit < itemService.ItemsCount() {
		newCursor := cursor + limit
		button, err := makeButton("К следующей странице", newCursor)
		if err != nil {
			return "", nil, err
		}
		buttons = append(buttons, *button)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons)
	return formatItems(items), &keyboard, nil
}

func makeButton(buttonText string, buttonCursor uint64) (*tgbotapi.InlineKeyboardButton, error) {
	offsetData := CursorData{
		Cursor: buttonCursor,
	}
	serCursorData, err := json.Marshal(offsetData)
	if err != nil {
		return nil, err
	}
	callbackPath := path.CallbackPath{
		Domain:       "product",
		Subdomain:    "item",
		CallbackName: "list",
		CallbackData: string(serCursorData),
	}
	button := tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackPath.String())
	return &button, nil
}

func formatItems(items []item.Item) string {
	if len(items) == 0 {
		return "Ни одного элемента"
	}
	var res strings.Builder
	for _, curItem := range items {
		res.WriteString(curItem.String() + ";\n")
	}
	return res.String()
}

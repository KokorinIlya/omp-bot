package item

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
)

//goland:noinspection GoNameStartsWithPackageName
type ItemService interface {
	Describe(itemId uint64) (*item.Item, error)
	List(cursor uint64, limit uint64) ([]item.Item, error)
	Create(item item.Item) (uint64, error)
	Update(itemId uint64, item item.Item) error
	Remove(itemId uint64) error
}

//goland:noinspection GoNameStartsWithPackageName
type ItemCommander struct {
	botApi      *tgbotapi.BotAPI
	itemService ItemService
}

func NewItemCommander(botApi *tgbotapi.BotAPI) *ItemCommander {
	service := item.NewDummyItemService()
	return &ItemCommander{
		botApi:      botApi,
		itemService: service,
	}
}

func (commander *ItemCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	default:
		log.Printf("ItemCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (commander *ItemCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		commander.Help(msg)
	case "get":
		commander.Get(msg)
	case "new":
		commander.New(msg)
	default:
		commander.Default(msg)
	}
}

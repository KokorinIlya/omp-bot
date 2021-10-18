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
	ItemsCount() uint64
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
	case "list":
		commander.CallbackList(callback, callbackPath)
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
	case "edit":
		commander.Edit(msg)
	case "delete":
		commander.Delete(msg)
	case "list":
		commander.List(msg)
	default:
		commander.Default(msg)
	}
}

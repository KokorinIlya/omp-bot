package product

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

//goland:noinspection GoNameStartsWithPackageName
type ProductCommander struct {
	botApi        *tgbotapi.BotAPI
	itemCommander Commander
}

func (commander *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "item":
		commander.itemCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ProductCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (commander *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "item":
		commander.itemCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ProductCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

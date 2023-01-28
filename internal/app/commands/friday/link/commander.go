package link

import (
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/friday/link"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type FridayLinkCommander struct {
	bot         *tgbotapi.BotAPI
	linkService *link.Service
}

func NewFridayLinkCommander(
	bot *tgbotapi.BotAPI,
) *FridayLinkCommander {
	linkService := link.NewService()

	return &FridayLinkCommander{
		bot:         bot,
		linkService: linkService,
	}
}

func (c *FridayLinkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("FridayLinkCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *FridayLinkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}

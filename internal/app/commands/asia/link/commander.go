package link

import (
	"github.com/ozonmp/omp-bot/internal/service/asia/link"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type AsiaLinkCommander struct {
	bot         *tgbotapi.BotAPI
	linkService *link.Service
}

func NewAsiaLinkCommander(
	bot *tgbotapi.BotAPI,
) *AsiaLinkCommander {
	linkService := link.NewService()

	return &AsiaLinkCommander{
		bot:         bot,
		linkService: linkService,
	}
}

func (c *AsiaLinkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("AsiaLinkCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *AsiaLinkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

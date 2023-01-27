package link

import (
	"github.com/ozonmp/omp-bot/internal/service/daily/link"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type DailyLinkCommander struct {
	bot         *tgbotapi.BotAPI
	linkService *link.Service
}

func NewDailyLinkCommander(
	bot *tgbotapi.BotAPI,
) *DailyLinkCommander {
	linkService := link.NewService()

	return &DailyLinkCommander{
		bot:         bot,
		linkService: linkService,
	}
}

func (c *DailyLinkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DailyLinkCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DailyLinkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

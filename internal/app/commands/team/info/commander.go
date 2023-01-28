package info

import (
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/team/info"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type TeamInfoCommander struct {
	bot         *tgbotapi.BotAPI
	infoService *info.Service
}

func NewTeamInfoCommander(
	bot *tgbotapi.BotAPI,
) *TeamInfoCommander {
	infoService := info.NewService()

	return &TeamInfoCommander{
		bot:         bot,
		infoService: infoService,
	}
}

func (c *TeamInfoCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoInfoCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *TeamInfoCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

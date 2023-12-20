package cdek

import (
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/check/cdek"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CheckCDEKCommander struct {
	bot         *tgbotapi.BotAPI
	cdekService *cdek.Service
}

func NewCheckPRFCommander(
	bot *tgbotapi.BotAPI,
) *CheckCDEKCommander {
	cdekService := cdek.NewService()

	return &CheckCDEKCommander{
		bot:         bot,
		cdekService: cdekService,
	}
}

func (c *CheckCDEKCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "help":
		c.CallbackHelp(callback)
	default:
		log.Printf("CheckCDEKCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CheckCDEKCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "start":
		c.Help(msg)
	case "help":
		c.Help(msg)
	case "by_id":
		c.Check(msg)
	default:
		log.Printf("CheckCDEKCommander.HandleCommand: unknown command - %s", commandPath.Subdomain)
	}
}

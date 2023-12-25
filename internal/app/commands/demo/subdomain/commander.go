package subdomain

import (
	"github.com/MostofWebmaker/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type DemoSubdomainCommander struct {
	bot *tgbotapi.BotAPI
}

func NewDemoSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *DemoSubdomainCommander {
	return &DemoSubdomainCommander{
		bot: bot,
	}
}

func (c *DemoSubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	log.Printf("Method is not implemented!")
}

func (c *DemoSubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "start":
		c.Help(msg)
	case "help":
		c.Help(msg)

	default:
		c.Default(msg)
	}
}

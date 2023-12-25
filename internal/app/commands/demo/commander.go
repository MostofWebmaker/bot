package demo

import (
	"github.com/MostofWebmaker/bot/internal/app/commands/demo/subdomain"
	"github.com/MostofWebmaker/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type DemoCommander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander Commander
}

func NewDemoCommander(
	bot *tgbotapi.BotAPI,
) *DemoCommander {
	return &DemoCommander{
		bot:                bot,
		subdomainCommander: subdomain.NewDemoSubdomainCommander(bot),
	}
}

func (c *DemoCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	log.Printf("Method is not implemented!")
}

func (c *DemoCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

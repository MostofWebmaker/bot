package daily

import (
	"github.com/MostofWebmaker/bot/internal/app/commands/daily/link"
	"github.com/MostofWebmaker/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LinkCommander struct {
	bot           *tgbotapi.BotAPI
	linkCommander Commander
}

func NewDailyLinkCommander(
	bot *tgbotapi.BotAPI,
) *LinkCommander {
	return &LinkCommander{
		bot: bot,
		// subdomainCommander
		linkCommander: link.NewDailyLinkCommander(bot),
	}
}

func (c *LinkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "link":
		c.linkCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LinkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "link":
		c.linkCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

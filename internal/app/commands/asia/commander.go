package asia

import (
	"github.com/MostofWebmaker/bot/internal/app/commands/asia/link"
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
	asiaCommander Commander
}

func NewAsiaLinkCommander(
	bot *tgbotapi.BotAPI,
) *LinkCommander {
	return &LinkCommander{
		bot: bot,
		// subdomainCommander
		asiaCommander: link.NewAsiaLinkCommander(bot),
	}
}

func (c *LinkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "link":
		c.asiaCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("AsiaLinkCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LinkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "link":
		c.asiaCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("AsiaLinkCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

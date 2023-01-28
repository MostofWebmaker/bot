package friday

import (
	"github.com/MostofWebmaker/bot/internal/app/commands/friday/link"
	"github.com/MostofWebmaker/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LinkCommander struct {
	bot             *tgbotapi.BotAPI
	fridayCommander Commander
}

func NewFridayLinkCommander(
	bot *tgbotapi.BotAPI,
) *LinkCommander {
	return &LinkCommander{
		bot: bot,
		// subdomainCommander
		fridayCommander: link.NewFridayLinkCommander(bot),
	}
}

func (c *LinkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "link":
		c.fridayCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("FridayLinkCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LinkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "link":
		c.fridayCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("FridayLinkCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

package friday

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/friday/link"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
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

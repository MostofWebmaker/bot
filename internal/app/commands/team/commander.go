package team

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/team/info"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type InfoCommander struct {
	bot           *tgbotapi.BotAPI
	infoCommander Commander
}

func NewTeamInfoCommander(
	bot *tgbotapi.BotAPI,
) *InfoCommander {
	return &InfoCommander{
		bot: bot,
		// subdomainCommander
		infoCommander: info.NewTeamInfoCommander(bot),
	}
}

func (c *InfoCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "info":
		c.infoCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *InfoCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "info":
		c.infoCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

package team

import (
	"github.com/MostofWebmaker/bot/internal/app/commands/team/info"
	"github.com/MostofWebmaker/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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

package check

import (
	"github.com/MostofWebmaker/bot/internal/app/commands/check/box"
	"github.com/MostofWebmaker/bot/internal/app/commands/check/kce"
	"github.com/MostofWebmaker/bot/internal/app/commands/check/prf"
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/cache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CheckCommander struct {
	bot           *tgbotapi.BotAPI
	c             *cache.Cache
	KCECommander  Commander
	PRFCommander  Commander
	CDEKCommander Commander
	BOXCommander  Commander
}

func NewCheckCommander(
	bot *tgbotapi.BotAPI,
	c *cache.Cache,
) *CheckCommander {
	return &CheckCommander{
		bot:          bot,
		KCECommander: kce.NewCheckKCECommander(bot, c),
		PRFCommander: prf.NewCheckPRFCommander(bot, c),
		BOXCommander: box.NewCheckBoxberryCommander(bot, c),
	}
}

func (c *CheckCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "kce":
		c.KCECommander.HandleCallback(callback, callbackPath)
	case "prf":
		c.PRFCommander.HandleCallback(callback, callbackPath)
	case "cdek":
		c.CDEKCommander.HandleCallback(callback, callbackPath)
	case "box":
		c.BOXCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CheckCommander.HandleCallback: unknown company_name - %s", callbackPath.Subdomain)
	}
}

func (c *CheckCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "kce":
		c.KCECommander.HandleCommand(msg, commandPath)
	case "prf":
		c.PRFCommander.HandleCommand(msg, commandPath)
	case "cdek":
		c.CDEKCommander.HandleCommand(msg, commandPath)
	case "box":
		c.BOXCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CheckCommander.HandleCommand: unknown company_name - %s", commandPath.Subdomain)
	}
}

package prf

import (
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/cache"
	"github.com/MostofWebmaker/bot/internal/service/check/prf"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const cacheKeyPrefix = "prf_"

type CheckPRFCommander struct {
	bot        *tgbotapi.BotAPI
	prfService *prf.Service
	c          *cache.Cache
}

func NewCheckPRFCommander(
	bot *tgbotapi.BotAPI,
	c *cache.Cache,
) *CheckPRFCommander {
	prfService := prf.NewService()

	return &CheckPRFCommander{
		bot:        bot,
		prfService: prfService,
		c:          c,
	}
}

func (c *CheckPRFCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "help":
		c.CallbackHelp(callback)
	default:
		log.Printf("CheckPRFCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CheckPRFCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "start":
		c.Help(msg)
	case "help":
		c.Help(msg)
	case "by_id":
		var args string
		if msg.CommandArguments() != "" {
			args = msg.CommandArguments()
		} else if commandPath.Args != "" {
			args = commandPath.Args
		}
		c.Check(args, msg.Chat.ID)
	default:
		log.Printf("PRFCommander.HandleCommand: unknown command - %s", commandPath.Subdomain)
	}
}

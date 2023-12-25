package kce

import (
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/cache"
	"github.com/MostofWebmaker/bot/internal/service/check/kce"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const cacheKeyPrefix = "kce_"

type CheckKCECommander struct {
	bot        *tgbotapi.BotAPI
	kceService *kce.Service
	c          *cache.Cache
}

func NewCheckKCECommander(
	bot *tgbotapi.BotAPI,
	c *cache.Cache,
) *CheckKCECommander {
	kceService := kce.NewService()

	return &CheckKCECommander{
		bot:        bot,
		kceService: kceService,
		c:          c,
	}
}

func (c *CheckKCECommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "help":
		c.CallbackHelp(callback)
	default:
		log.Printf("CheckKCECommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CheckKCECommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "start":
		c.Help(msg)
	case "help":
		c.Help(msg)
	case "by_id":
		c.Check(msg)
	default:
		log.Printf("KCECommander.HandleCommand: unknown command - %s", commandPath.Subdomain)
	}
}

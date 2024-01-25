package box

import (
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/cache"
	"github.com/MostofWebmaker/bot/internal/service/check/box"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const cacheKeyPrefix = "box_"

type CheckBoxberryCommander struct {
	bot        *tgbotapi.BotAPI
	boxService *box.Service
	c          *cache.Cache
}

func NewCheckBoxberryCommander(
	bot *tgbotapi.BotAPI,
	c *cache.Cache,
) *CheckBoxberryCommander {
	boxService := box.NewService()

	return &CheckBoxberryCommander{
		bot:        bot,
		boxService: boxService,
		c:          c,
	}
}

func (c *CheckBoxberryCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "help":
		c.CallbackHelp(callback)
	default:
		log.Printf("CheckBoxberryCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CheckBoxberryCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
		log.Printf("CheckBoxberryCommander.HandleCommand: unknown command - %s", commandPath.Subdomain)
	}
}

package router

import (
	"github.com/MostofWebmaker/bot/internal/app/commands/check"
	"github.com/MostofWebmaker/bot/internal/app/commands/demo"
	"github.com/MostofWebmaker/bot/internal/app/path"
	"github.com/MostofWebmaker/bot/internal/service/cache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	// bot
	bot            *tgbotapi.BotAPI
	demoCommander  Commander
	checkCommander Commander
}

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
	return &Router{
		// bot
		bot:            bot,
		demoCommander:  demo.NewDemoCommander(bot),
		checkCommander: check.NewCheckCommander(bot, cache.NewCache()),
	}
}

func (c *Router) CheckAccess(user *tgbotapi.User) bool {
	isBot := user.IsBot
	if isBot {
		return false
	}
	userName := user.UserName
	accessUsers, found := os.LookupEnv("ACCESS_USERS")
	if !found {
		log.Panic("environment variable ACCESS_USERS not found in .env")
	}

	accessList := strings.Split(accessUsers, ";")
	access := false
	for _, value := range accessList {
		if value == userName {
			access = true
		}
	}

	return access
}

func (c *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	var user interface{}
	if update.Message != nil {
		user = update.Message.From
	}

	if update.CallbackQuery != nil {
		user = update.CallbackQuery.From
	}

	if user == nil {
		log.Panic("cannot get user")
	}

	gotUser, ok := user.(*tgbotapi.User)
	if !ok {
		log.Panic("cannot convert user")
	}

	access := c.CheckAccess(gotUser)

	if !access {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"You don't have access to this bot! Go away, please.\n",
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("TeamInfoCommander.Help: error sending reply message to chat - %v", err)
		}

		return
	}

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "demo":
		c.demoCommander.HandleCallback(callback, callbackPath)
	case "check":
		c.checkCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	switch commandPath.Domain {
	case "demo":
		c.demoCommander.HandleCommand(msg, commandPath)
	case "help":
		c.demoCommander.HandleCommand(msg, commandPath)
	case "check":
		c.checkCommander.HandleCommand(msg, commandPath)
	//default:
	//	log.Printf("Router.handleCallback: unknown domain - %s", commandPath.Domain)
	default:
		c.demoCommander.HandleCommand(msg, commandPath)
	}
}

func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}_{company_name}_{track_no}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}

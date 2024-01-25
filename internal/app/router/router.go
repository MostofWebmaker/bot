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
	bot            *tgbotapi.BotAPI
	demoCommander  Commander
	checkCommander Commander
	c              *cache.Cache
}

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
	c := cache.NewCache()
	return &Router{
		c:              c,
		bot:            bot,
		demoCommander:  demo.NewDemoCommander(bot),
		checkCommander: check.NewCheckCommander(bot, c),
	}
}

func (c *Router) checkAccess(user *tgbotapi.User) bool {
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

	access := c.checkAccess(gotUser)

	if !access {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"У вас нет прав на пользование данным ботом, обратитесь к администратору.\n",
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
	var msgCommand, trackNo string
	if !msg.IsCommand() {
		cacheKey := msg.From.UserName + "check"
		value, _ := c.c.Get(cacheKey)
		if value != "" {
			c.c.Remove(cacheKey)
			trackNo = msg.Text
			msg.Text = value + " " + trackNo
			msgCommand = value[1:]
		} else {
			c.showCommandFormat(msg)

			return
		}
	} else {
		msgCommand = msg.Command()
	}

	commandPath, err := path.ParseCommand(msgCommand)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	if trackNo != "" {
		commandPath.Args = trackNo
	}

	switch commandPath.Domain {
	case "demo":
		c.demoCommander.HandleCommand(msg, commandPath)
	case "help":
		c.demoCommander.HandleCommand(msg, commandPath)
	case "check":
		c.checkCommander.HandleCommand(msg, commandPath)
	default:
		c.demoCommander.HandleCommand(msg, commandPath)
	}
}

func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{check}_{provider_name}_{by_id} {track_no}\n\n"+
		"Вместо {provider_name} подставляем: \n"+
		"Boxberry - box\n"+
		"KCE - kce\n"+
		"Почта России - prf\n\n"+
		"Вместо {track_no} подставляем трек номер заказа\n")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}

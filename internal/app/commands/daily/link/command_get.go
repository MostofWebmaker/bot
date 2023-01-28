package link

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *DailyLinkCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	stringArgs := strings.Split(args, " ")
	LinkType := stringArgs[0]
	product, err := c.linkService.Get(LinkType)
	if err != nil {
		log.Printf("fail to get product with LinkType %s: %v", LinkType, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Http,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DailyLinkCommander.Get: error sending reply message to chat - %v", err)
	}
}

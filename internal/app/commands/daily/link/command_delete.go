package link

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *DailyLinkCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	stringArgs := strings.Split(args, " ")
	linkType := stringArgs[0]

	//product, err := c.linkService.Get(LinkType)
	_, err := c.linkService.Delete(linkType)
	if err != nil {
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("successfully deleted link with linkType = %s", linkType),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DailyLinkCommander.Get: error sending reply message to chat - %v", err)
	}
}

package link

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *DailyLinkCommander) Get(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Daily link for you:\n"+
			"https://vcc.itbizstuff.com/b/ale-6a2-au9-adv\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("TeamInfoCommander.Help: error sending reply message to chat - %v", err)
	}
}

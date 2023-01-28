package link

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *AsiaLinkCommander) Get(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Asia link for you:\n"+
			"https://yandex.ru\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AsiaLinkCommander.Get: error sending reply message to chat - %v", err)
	}
}

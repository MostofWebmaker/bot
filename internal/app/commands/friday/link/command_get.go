package link

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *FridayLinkCommander) Get(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Weekly link for you:\n"+
			"https://yandex.ru\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("FridayLinkCommander.Get: error sending reply message to chat - %v", err)
	}
}

package subdomain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *DemoSubdomainCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__team__info - Team 15 info\n"+
			"/get__daily__link - get link of daily meeting\n"+
			"/get__friday__link - get link of results of week meeting\n"+
			"/get__asia__link - get link of results of asia meeting",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}

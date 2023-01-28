package link

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *FridayLinkCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Team 15 info:\n"+
			"developers:\n"+
			"- Demetrio Locatelli 12.10.1989\n"+
			"- Egorio Barnello \n"+
			"system analytic: Anastasiia B 04.12.1997\n"+
			"QA: Roman Empire 16.01.1991\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("FridayLinkCommander.Help: error sending reply message to chat - %v", err)
	}
}

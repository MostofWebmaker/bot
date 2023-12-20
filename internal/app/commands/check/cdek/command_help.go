package cdek

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *CheckCDEKCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Проверьте статус вашего заказа по одной из доступных транспортных компаний:\n"+
			"СДЭК:  /check_cdek {track_no}\n"+
			"KCE:  /check_kce {track_no}\n"+
			"Почта России:  /check_prf {track_no}\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckKCECommander.Help: error sending reply message to chat - %v", err)
	}
}

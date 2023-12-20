package subdomain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Boxberry", "check_box_help"),
		tgbotapi.NewInlineKeyboardButtonData("КСЕ", "check_kce_help"),
		tgbotapi.NewInlineKeyboardButtonData("Почта России", "check_prf_help"),
	),
)

func (c *DemoSubdomainCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Выберите одну из доступных транспортных компаний:\n",
	)
	msg.ReplyMarkup = numericKeyboard

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}

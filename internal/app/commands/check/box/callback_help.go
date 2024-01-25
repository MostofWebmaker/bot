package box

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CallbackHelpData struct {
	Offset int `json:"offset"`
}

func (c *CheckBoxberryCommander) CallbackHelp(callback *tgbotapi.CallbackQuery) {
	msgText := "Вы выбрали провайдера Boxberry для проверки статуса заказа: \n\n" +
		"Введите пожалуйста трек номер:\n"
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		msgText,
	)

	err := c.c.Set(callback.From.UserName+"check", "/check_box_by_id")
	if err != nil {
		log.Printf("CheckBoxberryCommander.CallbackHelp: error sending reply message to chat - %v", err)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckBoxberryCommander.CallbackHelp: error sending reply message to chat - %v", err)
	}
}

package kce

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CallbackHelpData struct {
	Offset int `json:"offset"`
}

func (c *CheckKCECommander) CallbackHelp(callback *tgbotapi.CallbackQuery) {
	msgText := "Вы выбрали провайдера KCE для проверки статуса заказа: \n\n" +
		"Введите пожалуйста трек номер заказа:\n"
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		msgText,
	)

	err := c.c.Set(callback.From.UserName+"check", "/check_kce_by_id")
	if err != nil {
		log.Printf("CheckKCECommander.CallbackHelp: error sending reply message to chat - %v", err)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckKCECommander.CallbackHelp: error sending reply message to chat - %v", err)
	}
}

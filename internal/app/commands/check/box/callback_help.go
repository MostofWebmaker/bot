package box

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CallbackHelpData struct {
	Offset int `json:"offset"`
}

func (c *CheckBoxberryCommander) CallbackHelp(callback *tgbotapi.CallbackQuery) {
	msgText := "Чтобы запросить статус по СДЭК выполните команду: \n" +
		"/check_box_by_id {track_no}\n\n" +
		"Вместо {track_no} необходимо вставить трек номер заказа."
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		msgText,
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckBoxberryCommander.CallbackHelp: error sending reply message to chat - %v", err)
	}
}

package kce

import (
	"encoding/json"
	"github.com/MostofWebmaker/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *CheckKCECommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products := c.kceService.List()
	for _, p := range products {
		outputMsgText += p.Data
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "check",
		Subdomain:    "kce",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckKCECommander.List: error sending reply message to chat - %v", err)
	}
}

package cdek

import (
	"encoding/json"
	"fmt"
	"github.com/MostofWebmaker/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *CheckCDEKCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("CheckKCECommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckKCECommander.CallbackList: error sending reply message to chat - %v", err)
	}
}

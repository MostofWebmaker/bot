package cdek

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *CheckCDEKCommander) Check(inputMessage *tgbotapi.Message) {
	preparedArgs := strings.Trim(inputMessage.CommandArguments(), " ")
	stringArgs := strings.Split(preparedArgs, " ")

	l := len(stringArgs)
	msgText := ""
	if l == 0 {
		msgText = "Track number is required! Please send with with track number."
	} else {
		trackNo := stringArgs[0]
		prfInfo, err := c.cdekService.Check(trackNo)
		if err != nil {
			log.Printf("fail to get product with trackNo %s: %v", trackNo, err)
			return
		}
		msgText = prfInfo.GetData()
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckKCECommander.Get: error sending reply message to chat - %v", err)
	}
}

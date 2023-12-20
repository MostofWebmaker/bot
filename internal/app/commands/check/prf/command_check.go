package prf

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *CheckPRFCommander) Check(inputMessage *tgbotapi.Message) {
	preparedArgs := strings.Trim(inputMessage.CommandArguments(), " ")
	stringArgs := strings.Split(preparedArgs, " ")

	l := len(stringArgs)
	msgText := ""
	if l == 0 {
		msgText = "Track number is required! Please send with with track number."
	} else {
		trackNo := stringArgs[0]
		data, err := c.c.Get(cacheKeyPrefix + trackNo)
		if err != nil {
			log.Printf("cannot get value by cache key %s. Need to do http request", cacheKeyPrefix+trackNo)
			prfInfo, err := c.prfService.Check(trackNo)
			if err != nil {
				log.Fatal("something went wrong")
			}
			err = c.c.Set(cacheKeyPrefix+trackNo, prfInfo.GetData())
			if err != nil {
				log.Fatal("got error on set cache value")
				return
			}
			data = prfInfo.GetData()
		} else {
			log.Printf("successful got value by cache key `%s`", cacheKeyPrefix+trackNo)
		}
		msgText = data
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

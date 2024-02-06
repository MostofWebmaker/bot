package prf

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *CheckPRFCommander) Check(args string, chatID int64) {
	preparedArgs := strings.Trim(args, " ")
	stringArgs := strings.Split(preparedArgs, " ")

	l := len(stringArgs)
	msgText := ""
	if l == 0 {
		msgText = "Трек номер обязателен для данной команды! Пожалуйста отправьте с трек номером."
	} else {
		trackNo := stringArgs[0]
		data, err := c.c.Get(cacheKeyPrefix + trackNo)
		if err != nil {
			log.Printf("cannot get value by cache key %s. Need to do http request", cacheKeyPrefix+trackNo)
			prfInfo, err := c.prfService.Check(trackNo)
			if err != nil {
				log.Printf("got error from prf provider %s", err.Error())
			}
			data, err = prfInfo.GetData()
			if err != nil {
				log.Printf("got error in parsing data from prf provider %s", err.Error())
				data = "К сожалению не удалось запросить данные от провайдера. Пожалуйста попробуйте еще раз позже."
			} else {
				err = c.c.Set(cacheKeyPrefix+trackNo, data)
			}
		} else {
			log.Printf("successful got value by cache key `%s`", cacheKeyPrefix+trackNo)
		}
		msgText = data
	}

	msg := tgbotapi.NewMessage(
		chatID,
		msgText,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CheckPRFCommander.Check: error sending reply message to chat - %v", err)
	}
}

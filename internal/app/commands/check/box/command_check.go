package box

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *CheckBoxberryCommander) Check(args string, chatID int64) {
	preparedArgs := strings.Trim(args, " ")
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
			boxInfo, err := c.boxService.Check(trackNo)
			if err != nil {
				data = err.Error()
			} else {
				data, err = boxInfo.GetData()
				if err != nil {
					data = err.Error()
				} else {
					err = c.c.Set(cacheKeyPrefix+trackNo, data)
				}
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
		log.Printf("CheckBoxberryCommander.Check: error sending reply message to chat - %v", err)
	}
}

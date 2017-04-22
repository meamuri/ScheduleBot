package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"

	"strconv"
)

func getInfo(date int) string {
	switch {

	// на ввод поступило слишком маленькое число
	case date < 24:
		return EarlyForFestival

	// все кейсы верного диапазона
	case date == 24:
		return InfoAboutDate24
	case date == 25:
		return InfoAboutDate25
	case date == 26:
		return InfoAboutDate26
	case date == 27:
		return InfoAboutDate27
	case date == 28:
		return InfoAboutDate28
	case date == 29:
		return InfoAboutDate29

	// слишком большое число!
	case date > 29:
		return TooLateForFestival
	// любой вид ошибки, хотя данный свитч и покрывает все кейсы
	default:
		return InfoAboutUnknownDate
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(APIToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		var answer string
		num, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			answer = InfoAboutNotDate
		} else {
			answer = getInfo(num)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}

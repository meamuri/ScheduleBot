package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"

	"strconv"
)

const (
	// InfoAboutDate24 is
	InfoAboutDate24 string = "Показы 24 апреля: \n-Шарли и его зубы \n-Скупость \n-Парикмахер \n-Светлая сторона \n-Новая жизнь \n-Пеле"

	// InfoAboutDate25 is
	InfoAboutDate25 string = "Показы 25 апреля: \n-Фриланс \n-Лето \n-Утопия \n-Муза"

	// InfoAboutUnknownDate is
	InfoAboutUnknownDate string = "Сожалеем, в этот день фестиваль не проводится!"

	// InfoAboutNotDate is
	InfoAboutNotDate string = "Пожалуйста, введите интересующий вас день показа (23 - 25)"
)

func getInfo(date int) string {
	switch date {
	case 24:
		return InfoAboutDate24
	case 25:
		return InfoAboutDate25
	default:
		return InfoAboutUnknownDate
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
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

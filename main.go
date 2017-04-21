package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"

	"strconv"
)

const (
	// EarlyForFestival - если пользователь ввел число, меньшее чем дата начала фестиваля
	EarlyForFestival string = "В этот день фестиваль еще не начался!"
	// ToLateForFestival - а эта строка о том, что введенное число больше даты последнего дня фестиваля
	ToLateForFestival string = "Сожалеем, фестиваль уже закончился"
	//InfoAboutDate24 - расписание на 24 апреля
	InfoAboutDate24 string = "Показы 24 апреля: \n-Шарли и его зубы \n-Скупость \n-Парикмахер \n-Светлая сторона \n-Новая жизнь \n-Пеле"
	// InfoAboutDate25 - расписание на 25 апреля
	InfoAboutDate25 string = "Показы 25 апреля: \n-Фриланс \n-Лето \n-Утопия \n-Муза"
	// InfoAboutUnknownDate - строка для любой ошибки (пусть формально такой кейс и невозможен в силу предикатов)
	InfoAboutUnknownDate string = "Сожалеем, в этот день фестиваль не проводится!"
	// InfoAboutNotDate - любая ошибка ввода
	InfoAboutNotDate string = "Пожалуйста, введите интересующий вас день показа (23 - 25)"
)

func getInfo(date int) string {
	switch {

	case date < 24:
		return EarlyForFestival

	case date == 24:
		return InfoAboutDate24
	case date == 25:
		return InfoAboutDate25

	case date > 29:
		return ToLateForFestival
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

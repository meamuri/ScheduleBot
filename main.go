package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"

	"strconv"
)

const (
	// EarlyForFestival - если пользователь ввел число, меньшее чем дата начала фестиваля
	EarlyForFestival string = "В этот день фестиваль еще не начался!"
	// TooLateForFestival - а эта строка о том, что введенное число больше даты последнего дня фестиваля
	TooLateForFestival string = "Сожалеем, фестиваль уже закончился"

	// PrefaceText - текст с которым конкатенируются все последующие строки
	PrefaceText string = "Когда: с 19:00 до 21:00\n" +
		"Где: 	кинотеатр спартак\n" +
		"Что:	показы "

	// InfoAboutDate24 - расписание на 24 апреля
	InfoAboutDate24 string = PrefaceText + "24 апреля\n" +
		"-Шарли и его зубы\n" +
		"-Скупость\n" +
		"-Парикмахер\n" +
		"-Светлая сторона\n" +
		"-Новая жизнь\n" +
		"-Пеле\n" +
		"-История о толстой овечке\n" +
		"-Ука\n" +
		"-Апорт\n" +
		"-Бельчонок и санки\n" +
		"-Восхождение\n" +
		"-Два трамвая\n" +
		"-Какие сны снятся медведю\n" +
		"-Мальчук-кот и новый год\n" +
		"-Мой робот\n" +
		"-ПинWin\n" +
		"-Письма куклы\n" +
		"-Тучка и кит\n" +
		"-Кукушка\n" +
		"-Оставить след\n" +
		"-Тип-топ\n" +
		"-Притягательный след\n" +
		"-Положитесь на меня\n" +
		"-Завершенный\n" +
		"-Флипоскоп\n" +
		"-Без жизни\n" +
		"-Хуан и облачко"

	// InfoAboutDate25 - расписание на 25 апреля
	InfoAboutDate25 string = PrefaceText + "25 апреля:\n" +
		"-Фриланс\n" +
		"-Лето\n" +
		"-Утопия\n" +
		"-Закон исключенного третьего, или Третьего не дано\n" +
		"-Муза\n" +
		"-Бьянка\n" +
		"-Заложница\n" +
		"-История юноши с волшебными веснушками\n" +
		"-Я не одинок"

	// InfoAboutDate26 - расписание на 26 апреля
	InfoAboutDate26 string = PrefaceText + "26 апреля:\n" +
		"-Сожалею\n" +
		"-Женёк\n" +
		"-Строгий выгоовор\n" +
		"-У меня есть брат\n" +
		"-Марджи\n" +
		"-Каждый 88\n" +
		"-Наперегонки"

	// InfoAboutDate27 - расписание на 27 апреля
	InfoAboutDate27 string = PrefaceText + "27 апреля:\n" +
		"-Мой дедушка был вишней\n" +
		"-Все в ящик\n" +
		"-Мама\n" +
		"-Не отпускай\n" +
		"-Слепая юность\n" +
		"-Кукла\n" +
		"-Друзья по переписке\n" +
		"-Добро пожаловать в чудесный мир\n" +
		"-Рубеж"

	// InfoAboutDate28 - расписание на 28 апреля
	InfoAboutDate28 string = PrefaceText + "28 апреля:\n" +
		"-Даррел\n" +
		"-Шоколадный ветер\n" +
		"-Работа мечты\n" +
		"-#КОГОУБИВАТЬ\n" +
		"-Запретная жизнь\n" +
		"-Цирк\n" +
		"-Фотографии\n" +
		"-Тузик, служить!\n" +
		"-Двойник\n" +
		"-Мягкая карамель\n" +
		"-Лучший город на земле"

	// InfoAboutDate29 - расписание на 29 апреля
	InfoAboutDate29 string = "29 апреля с 18:00 до 20:00 в нашем любимом кинотеатре" +
		"\"Спартак\" состоится итоговый показ фильмов победителей" +
		"кинофестиваля \"Новый горизонт\" - 2017! Обязательно ждем вас!"

	// InfoAboutUnknownDate - строка для любой ошибки (пусть формально такой кейс и невозможен в силу предикатов)
	InfoAboutUnknownDate string = "Сожалеем, в этот день фестиваль не проводится!"

	// InfoAboutNotDate - любая ошибка ввода
	InfoAboutNotDate string = "Пожалуйста, введите интересующий вас день показа (23 - 25)"
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

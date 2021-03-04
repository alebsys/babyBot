package telegram

import (
	"fmt"
	"log"

	"github.com/alebsys/baby-bot/config"
	"github.com/alebsys/baby-bot/pkg/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

// TODO не вынести ли в handlers.go?
// Weight структура для ввода и получения данных
type Weight struct {
	Date   string
	Weight float64
	ID     int
}

var (
	weight Weight

	menu = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	get  = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	back = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	// menu buttons.
	btnPostValue = menu.Text("Ввести свой вес")
	btnGetMenu   = menu.Text("Посмотреть статистику")

	// get buttons
	btnGetGraph = menu.Text("Построить график")
	btnGetDate  = menu.Text("Узнать вес за определенную дату")

	// back button
	btnBackMenu = menu.Text("Вернуться в меню")
)

// NewBot ...
func NewBot(configBot tb.Settings) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Println("Hello, I am babyBot!")

	collection := db.InitCollection()
	b, err := tb.NewBot(configBot)
	if err != nil {
		log.Fatal(err)
	}

	// Активирует стартовое меню бота
	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Привет!", menu)
	})

	// Главное меню
	menu.Reply(
		menu.Row(btnPostValue),
		menu.Row(btnGetMenu),
	)

	// Меню получения статистики
	get.Reply(
		menu.Row(btnGetGraph, btnGetDate),
		menu.Row(btnBackMenu),
	)

	// Меню кнопки возврата
	back.Reply(
		menu.Row(btnBackMenu),
	)

	// Обрабатывает ввод данных
	b.Handle(&btnPostValue, func(m *tb.Message) {
		b.Send(m.Sender, "Введите дату и свой вес в кг. Пример: `21/10/21 80.3`.", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			if err := generateValue(m, b, &weight); err != nil {
				return
			}
			postValue(m, collection, b, weight)
			b.Send(m.Sender, "Значение добавлено в базу данных.", menu)
		})
	})

	// Вход в меню получения статистики
	b.Handle(&btnGetMenu, func(m *tb.Message) {
		weight = Weight{}
		b.Send(m.Sender, "Что вы хотите посмотреть?", get)
	})

	// Обрабатывает получение данных за определенную дату
	b.Handle(&btnGetDate, func(m *tb.Message) {
		weight = Weight{}
		b.Send(m.Sender, "Введите интересующую вас дату. Пример: `21/10/21`.", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			if err := generateDate(m, b, &weight); err != nil {
				return
			}
			getDate(m, collection, b, weight)
		})
	})

	// Обрабатывает генерацию графика
	b.Handle(&btnGetGraph, func(m *tb.Message) {
		getGraph(m, collection, b)

	})

	// Выход в стартовое меню
	b.Handle(&btnBackMenu, func(m *tb.Message) {
		b.Send(m.Sender, "Давайте заново!", menu)
	})

	b.Start()
}

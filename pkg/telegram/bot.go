package telegram

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/alebsys/baby-bot/config"
	"github.com/alebsys/baby-bot/pkg/db"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Weight структура для ввода и получения данных
type Weight struct {
	Date   string
	Weight float64
	ID     int
}

var (
	weight Weight
	// B telebot
	B          *tb.Bot
	collection *mongo.Collection

	//m *tb.Message

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

func init() {

	log.SetFormatter(&log.JSONFormatter{})

	if err := config.Init(); err != nil {
		log.Fatal("error from config.Init(): ", err)
	}
	collection = db.InitCollection()

	// create bot
	var err error
	B, err = tb.NewBot(tb.Settings{
		URL:    "",
		Token:  viper.GetString("apiToken"),
		Poller: &tb.LongPoller{Timeout: viper.GetDuration("pollerTimeout") * time.Second},
	})
	if err != nil {
		log.Fatal("error from NewBot: ", err)
	}
}

// StartBot ...
func StartBot() {
	fmt.Println("Hello, I am babyBot!")

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

	// Активирует стартовое меню бота
	B.Handle("/start", start)

	// Обрабатывает ввод данных
	B.Handle(&btnPostValue, postMenu)

	// Вход в меню получения статистики
	B.Handle(&btnGetMenu, getMenu)

	// Обрабатывает получение данных за определенную дату
	B.Handle(&btnGetDate, getMenuDate)

	// Обрабатывает генерацию графика
	B.Handle(&btnGetGraph, getMenuGraph)

	// Выход в стартовое меню
	B.Handle(&btnBackMenu, backMenu)

	B.Start()
}

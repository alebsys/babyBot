package telegram

import (
	"fmt"
	"log"
	//"time"
	"github.com/alebsys/baby-bot/config"
	"github.com/alebsys/baby-bot/pkg/db"

	//"github.com/spf13/viper"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Weight ...
type Weight struct {
	Date   string
	Weight float64
	ID     int
}

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

	b.Handle("/post", func(m *tb.Message) {
		postHandler(m, collection)
	})

	b.Handle("/get", func(m *tb.Message) {
		getHandler(m, collection, b)
	})

	b.Handle("/delete", func(m *tb.Message) {
		deleteHandler(m, collection, b)
	})

	b.Handle("/graph", func(m *tb.Message) {
		graphHandler(m, collection, b)
	})

	b.Start()
}

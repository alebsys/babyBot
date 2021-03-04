package main

import (
	"log"
	"time"

	"github.com/alebsys/baby-bot/config"
	bot "github.com/alebsys/baby-bot/pkg/telegram"
	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	configBot := tb.Settings{
		URL:    "",
		Token:  viper.GetString("apiToken"),
		Poller: &tb.LongPoller{Timeout: viper.GetDuration("pollerTimeout") * time.Second},
	}
	bot.NewBot(configBot)
}

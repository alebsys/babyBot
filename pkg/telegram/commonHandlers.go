package telegram

import (
	"context"
	"errors"
	"fmt"
	l "github.com/alebsys/baby-bot/pkg/logs"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"

	"go.mongodb.org/mongo-driver/bson"
)

// start ...
func start(m *tb.Message) {
	if !m.Private() {
		return
	}
	if _, err := B.Send(m.Sender, "Привет!", menu); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}
	B.Handle(tb.OnText, func(m *tb.Message) {
		if _, err := B.Send(m.Sender, "Выберите один из пунктов меню.", menu); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
	})
}

// TODO создать функцию для удаления данных
func deleteHandler(m *tb.Message, collection *mongo.Collection, b *tb.Bot) {
	valueToSlice := strings.Split(m.Text, " ")
	dateValue := valueToSlice[1]

	filter := bson.D{{Key: "date", Value: dateValue}, {Key: "id", Value: m.Sender.ID}}

	// response, _ := fmt.Printf("Данные за %v удалены\n", sliceToInt)

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		//fmt.Println(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	_, err = b.Send(m.Sender, "Данные за "+dateValue+" удалены")
	if err != nil {
		log.Fatal(err)
	}
}

// validationDate validation of the entered date
func validationDate(s string) error {
	d, err := time.Parse("02/01/06", s)
	if err != nil {
		return errors.New("problem parsing date")
	}
	if time.Now().Before(d) {
		return errors.New("error from future")
	}
	return nil
}

// validationWeight validation of the entered weight
func validationWeight(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return errors.New("Error from validationWeight")
	}
	return nil
}

func backMenu(m *tb.Message) {
	if _, err := B.Send(m.Sender, "Давайте заново!", menu); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}
	B.Handle(tb.OnText, func(m *tb.Message) {
		if _, err := B.Send(m.Sender, "Выберите один из пунктов меню.", menu); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
	})
}

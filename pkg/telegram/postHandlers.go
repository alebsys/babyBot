package telegram

import (
	"context"
	"errors"
	l "github.com/alebsys/baby-bot/pkg/logs"
	"go.uber.org/zap"

	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	postMenuMessage = "Введите дату `число/месяц/год` и вес в `кг`.\n\n" +
		"Примеры:\n- `21/10/20 80.3`\n- `01/10/20 65`"

	wrongDateInputMessage = "*Неверный формат даты!*\n\nДата должна быть:\n" +
		"1. в формате `число/месяц/год`\n" +
		"2. за сегодняшнее или предыдущие числа\n\nПример: `21/10/20`"

	wrongWeightInputMessage = "*Неверный формат значения веса*!\n\nВес должен быть:\n" +
		"1. в формате десятичного числа\n" +
		"2. с разделителем в виде точки (`.`)\n\n" +
		"Пример `80` или `76.6`"
)

var (
	markdownOn = &tb.SendOptions{ParseMode: tb.ModeMarkdown,}
)

func postMenu(m *tb.Message) {

	weight = Weight{}

	if _, err := B.Send(m.Sender, postMenuMessage, markdownOn, back); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}

	B.Handle(tb.OnText, func(m *tb.Message) {
		if err := generateValue(m, B, &weight); err != nil {
			return
		}
		postValue(collection, weight, m)
		if _, err := B.Send(m.Sender, "Значение добавлено в базу данных.", menu); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}

		B.Handle(tb.OnText, func(m *tb.Message) {
			if _, err := B.Send(m.Sender, "Выберите один из пунктов меню.", menu); err != nil {
				l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
			}
		})
	})
}

// generateValue подготавливает дату и значение веса для записи в БД
func generateValue(m *tb.Message, b *tb.Bot, weight *Weight) error {
	s := strings.Split(m.Text, " ")

	if err := validationDate(s[0]); err != nil {
		if _, err := b.Send(m.Sender, wrongDateInputMessage, markdownOn); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).
			Infof("the user entered the data incorrectly: %v", m.Text)
		return errors.New("error from validationDate")
	}

	if err := validationWeight(s[1]); err != nil {
		if _, err = b.Send(m.Sender, wrongWeightInputMessage, markdownOn); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).
			Infof("the user entered the data incorrectly: %v", m.Text)
		return errors.New("error from validationWeight")
	}

	weight.Date = s[0]
	weight.Weight, _ = strconv.ParseFloat(s[1], 64)
	weight.ID = m.Sender.ID

	return nil
}

// postValue записывает данные в БД исходя из созданных в generateValue
func postValue(c *mongo.Collection, w Weight, m *tb.Message) {

	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{Key: "date", Value: w.Date}, {Key: "id", Value: w.ID}}
	_ = c.FindOne(context.TODO(), filter).Decode(&find)

	// Если не находим, то создаём запись в БД
	if find.Weight == 0 {
		if _, err := c.InsertOne(context.TODO(), w); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}

		// Иначе обновляем значение
	} else {
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "weight", Value: w.Weight},
			}},
		}
		if _, err := c.UpdateOne(context.TODO(), filter, update); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
	}
	l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).
		Infof("Value insert in the database: %v", m.Text)
}

package telegram

import (
	"context"
	"errors"
	//"log"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
	//"github.com/alebsys/baby-bot/pkg/logr"
	log "github.com/sirupsen/logrus"
)

var (
	postMenuMessage = "Введите дату `число/месяц/год` и вес в `кг`.\n\n" +
		"Примеры:\n- `21/10/20 80.3`\n- `01/10/20 65`"

	wrongDateInputMessage = "*Неверный формат даты!*\n\nДата должна быть:\n" +
		"1. в формате `число/месяц/год`\n" +
		"2. за сегодняшнее или предыдущие числа\n\nПример: `21/10/20`"

	wrongWeightInputMessage = "*Неверный формат значения веса*!\n\nВес должен быть:\n" +
		"1. в формате десятичного числа\n" +
		"2. с разделителем в виде точки (`.`)\n\n" +
		"Пример `80` или `76.6`"

	markdownOn = &tb.SendOptions{
		ParseMode: tb.ModeMarkdown,
	}
)

func postMenu(m *tb.Message) {

	weight = Weight{}

	_, _ = B.Send(m.Sender, postMenuMessage, markdownOn, back)

	B.Handle(tb.OnText, func(m *tb.Message) {
		if err := generateValue(m, B, &weight); err != nil {
			return
		}
		postValue(collection, weight, m)
		_, _ = B.Send(m.Sender, "Значение добавлено в базу данных.", menu)

		B.Handle(tb.OnText, func(m *tb.Message) {
			_, _ = B.Send(m.Sender, "Выберите один из пунктов меню.", menu)
		})
	})
}

// generateValue подготавливает дату и значение веса для записи в БД
func generateValue(m *tb.Message, b *tb.Bot, weight *Weight) error {
	logger := log.WithFields(log.Fields{
		"clientID":  m.Sender.ID,
		"messageID": m.ID,
		"function":  "generateValue",
	})
	s := strings.Split(m.Text, " ")

	if err := validationDate(s[0]); err != nil {
		_, _ = b.Send(m.Sender, wrongDateInputMessage, markdownOn)
		logger.Info("the user entered the data incorrectly - the date ")
		return errors.New("error from validationDate")
	}

	if err := validationWeight(s[1]); err != nil {
		_, _ = b.Send(m.Sender, wrongWeightInputMessage, markdownOn)
		logger.Info("the user entered the data incorrectly - the weight ")
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
	logger := log.WithFields(log.Fields{
		"clientID":  m.Sender.ID,
		"messageID": m.ID,
		"function":  "postValue",
	})

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{Key: "date", Value: w.Date}, {Key: "id", Value: w.ID}}
	_ = c.FindOne(context.TODO(), filter).Decode(&find)

	// Если не находим, то создаём запись в БД
	if find.Weight == 0 {
		_, err := c.InsertOne(context.TODO(), w)
		if err != nil {
			logger.Error("failed to update the value in the database", err)
		}

		// Иначе обновляем значение
	} else {
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "weight", Value: w.Weight},
			}},
		}
		_, err := c.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			logger.Error("failed to update the value in the database", err)
		}
	}
	logger.Info("Value insert in the database")
}

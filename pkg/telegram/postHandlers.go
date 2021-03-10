package telegram

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

func postMenu(m *tb.Message) {
	weight = Weight{}
	B.Send(m.Sender, "Введите дату(число/месяц/год) и свой вес в кг.\nПример: `21/10/20 80.3` или `01/10/20 65`.", back)
	B.Handle(tb.OnText, func(m *tb.Message) {
		if err := generateValue(m, B, &weight); err != nil {
			return
		}
		postValue(m, collection, B, weight)
		B.Send(m.Sender, "Значение добавлено в базу данных.", menu)
		B.Handle(tb.OnText, func(m *tb.Message) {
			B.Send(m.Sender, "Выберите один из пунктов меню.", menu)
		})
	})
}

// generateValue подготавливает дату и значение веса для записи в БД
func generateValue(m *tb.Message, b *tb.Bot, weight *Weight) error {
	s := strings.Split(m.Text, " ")

	// if err := validationDate(s, b, m); err != nil {
	if err := validationDate(s[0]); err != nil {

		b.Send(m.Sender, "Неверный формат даты!\nДата должна быть:\n* в формате *число/месяц/год*\n* за сегодняшнее или предыдущие числа\nПример `21/10/21`")
		return errors.New("Error from validationDate")
	}

	if err := validationWeight(s[1]); err != nil {

		b.Send(m.Sender, "Неверный формат значения веса!\nВес должен быть:\n* в формате десятичного числа\n* с разделителем в виде точки (.)\nПример `80` или `76.6`")
		return errors.New("Error from validationWeight")
	}
	// After the checks carried out, we assign the values to the variable
	weight.Date = s[0]
	weight.Weight, _ = strconv.ParseFloat(s[1], 64)
	weight.ID = m.Sender.ID

	return nil
}

// postValue записывает данные в БД исходя из созданных в generateValue
func postValue(m *tb.Message, c *mongo.Collection, b *tb.Bot, w Weight) {

	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{Key: "date", Value: w.Date}, {Key: "id", Value: w.ID}}
	_ = c.FindOne(context.TODO(), filter).Decode(&find)

	// Если не находим, то создаём запись в БД
	if find.Weight == 0 {
		_, err := c.InsertOne(context.TODO(), w)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Значение добавлено.")

		// Иначе обновляем значение
	} else {
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "weight", Value: w.Weight},
			}},
		}
		_, err := c.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Значение обновлено.")
	}
}

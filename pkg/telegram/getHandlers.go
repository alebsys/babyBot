package telegram

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/wcharczuk/go-chart/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"
)

// getMenu ...
func getMenu(m *tb.Message) {
	B.Send(m.Sender, "Что вы хотите посмотреть?", get)
	B.Handle(tb.OnText, func(m *tb.Message) {
		B.Send(m.Sender, "Выберите один из пунктов меню.", get)
	})
}

// getMenuDate ...
func getMenuDate(m *tb.Message) {
	B.Send(m.Sender, "Введите интересующую вас дату (число/месяц/год).\nПример: `21/10/20`.", back)
	B.Handle(tb.OnText, func(m *tb.Message) {
		if err := generateDate(m, B, &weight); err != nil {
			return
		}
		if err := getDate(m, collection, B, weight); err != nil {
			return
		}
	})
}

// generateDate подготавливает структуру данных для получения данных из БД
func generateDate(m *tb.Message, b *tb.Bot, weight *Weight) error {
	s := strings.Split(m.Text, " ")

	if err := validationDate(s[0]); err != nil {
		b.Send(m.Sender, "Неверный формат даты!\nДата должна быть:\n* в формате *число/месяц/год*\n* за сегодняшнее или предыдущие числа\nПример `21/10/21`")
		return errors.New("error from generateDate")
	}
	weight.Date = m.Text
	return nil
}

// getDate получает данные из БД исходя из переданной даты
func getDate(m *tb.Message, collection *mongo.Collection, b *tb.Bot, weight Weight) error {
	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{Key: "date", Value: weight.Date}, {Key: "id", Value: m.Sender.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&find)

	if err != nil {
		log.Printf("event from collection.FindOne: %v", err)
		_, err := b.Send(m.Sender, "Данные за "+weight.Date+" отсутствуют.")
		if err != nil {
			log.Printf("error from b.Send if collection.FindOne err != nil : %v", err)
		}
		return errors.New("error from getDate")
	}
	_, err = b.Send(m.Sender, "Вес за "+find.Date+" -- "+fmt.Sprintf("%.1f", find.Weight)+"кг")
	if err != nil {
		log.Printf("error from b.Send if collection.FindOne err == nil : %v", err)
	}
	return nil
}

// TODO генерировать график исходя из дат по оси X
// getGraph генерирует график из введенных раннее данных
// getMenuGraph ...
func getMenuGraph(m *tb.Message) {
	filter := bson.D{{Key: "id", Value: m.Sender.ID}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("c.Find ERROR:", err)
	}

	var y []float64
	var yFloat []float64

	for cursor.Next(context.TODO()) {
		var p Weight

		// Decode the document
		if err := cursor.Decode(&p); err != nil {
			log.Fatal("cursor.Decode ERROR:", err)
		}

		y = append(y, p.Weight)

	}
	for i := range y {
		yFloat = append(yFloat, float64(i))
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: yFloat,
				YValues: y,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})

	err = graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Printf("graph.Render ERROR: %v", err)
		B.Send(m.Sender, "Слишком мало данных для отображения, минимум 2 записи.\nДобавьте еще!")
		return
	}

	p := &tb.Photo{File: tb.FromReader(buffer)}
	_, err = B.SendAlbum(m.Sender, tb.Album{p})
	if err != nil {
		log.Fatal("SendPhoto ERROR:", err)
	}
}

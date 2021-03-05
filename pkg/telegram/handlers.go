package telegram

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/wcharczuk/go-chart/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO создать функцию для удаления данных
func deleteHandler(m *tb.Message, collection *mongo.Collection, b *tb.Bot) {
	valueToSlice := strings.Split(m.Text, " ")
	dateValue := valueToSlice[1]

	filter := bson.D{{"date", dateValue}, {"id", m.Sender.ID}}

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

// TODO генерировать график исходя из дат по оси X
// getGraph генерирует график из введенных раннее данных
func getGraph(m *tb.Message, c *mongo.Collection, b *tb.Bot) {
	filter := bson.D{{"id", m.Sender.ID}}
	cursor, err := c.Find(context.TODO(), filter)
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
		b.Send(m.Sender, "Слишком мало данных для отображения. Добавьте еще!")
		return
	}

	p := &tb.Photo{File: tb.FromReader(buffer)}
	_, err = b.SendAlbum(m.Sender, tb.Album{p})
	if err != nil {
		log.Fatal("SendPhoto ERROR:", err)
	}
}

// generateDate подготавливает структуру данных для получения данных из БД
func generateDate(m *tb.Message, b *tb.Bot, weight *Weight) error {
	s := strings.Split(m.Text, " ")

	if err := validationDate(s, b); err != nil {
		b.Send(m.Sender, "Неверный формат даты!\nВведите дату (число/месяц/год) в формате `21/10/21 80.5`.")
		return errors.New("Error from generateValue")
	}
	weight.Date = m.Text
	return nil
}

// getDate получает данные из БД исходя из переданной даты
func getDate(m *tb.Message, collection *mongo.Collection, b *tb.Bot, weight Weight) error {
	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{"date", weight.Date}, {"id", m.Sender.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&find)

	if err != nil {
		fmt.Println("Нет таких данных")
		_, err := b.Send(m.Sender, "Данные за "+weight.Date+" отсутствуют.")
		if err != nil {
			log.Fatal(err)
		}
		return errors.New("Error from getDate")
	}
	_, err = b.Send(m.Sender, "Вес за "+find.Date+" -- "+fmt.Sprintf("%.1f", find.Weight)+"кг")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// generateValue подготавливает дату и значение веса для записи в БД
func generateValue(m *tb.Message, b *tb.Bot, weight *Weight) error {
	s := strings.Split(m.Text, " ")

	if err := validationDate(s, b); err != nil {
		b.Send(m.Sender, "Неверный формат даты!\nВведите дату (число/месяц/год) и вес (в кг) в формате `21/10/21 80.5`.")
		return errors.New("Error from generateValue")
	}

	if err := validationWeigth(s, b); err != nil {
		b.Send(m.Sender, "Неверный формат значения веса!\nВведите дату (число/месяц/год) и вес (в кг) в формате `21/10/21 80.5`.")
		return errors.New("Error from generateValue")
	}
	// After the checks carried out, we assign the values to the variable
	weight.Date = s[0]
	weight.Weight, _ = strconv.ParseFloat(s[1], 64)
	weight.ID = m.Sender.ID

	return nil
}

// validationDate validation of the entered date
func validationDate(s []string, b *tb.Bot) error {
	_, err := time.Parse("02/01/06", s[0])
	if err != nil {
		return errors.New("Error from validationDate")
	}
	return nil
}

// validationWeigth validation of the entered weight
func validationWeigth(s []string, b *tb.Bot) error {
	_, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return errors.New("Error from validationWeigth")
	}
	return nil
}

// postValue записывает данные в БД исходя из созданных в generateValue
func postValue(m *tb.Message, c *mongo.Collection, b *tb.Bot, w Weight) {

	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{"date", w.Date}, {"id", w.ID}}
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
			{"$set", bson.D{
				{"weight", w.Weight},
			}},
		}
		_, err := c.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Значение обновлено.")
	}
}

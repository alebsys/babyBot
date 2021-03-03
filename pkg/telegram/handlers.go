package telegram

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/wcharczuk/go-chart/v2"
	"go.mongodb.org/mongo-driver/bson"
)

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

func graphHandler(m *tb.Message, collection *mongo.Collection, b *tb.Bot) {
	filter := bson.D{{"id", m.Sender.ID}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("collection.Find ERROR:", err)
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

func postToday(m *tb.Message, collection *mongo.Collection, b *tb.Bot, dateValue string) {

	// Узнаем значение введеного веса
	valueToSlice := strings.Split(m.Text, " ")
	weightValue, err := strconv.ParseFloat(valueToSlice[0], 64)
	if err != nil {
		b.Send(m.Sender, "Введите корректное значение веса. Пример: 60.5")
	}
	postValue := Weight{Date: dateValue, Weight: weightValue, ID: m.Sender.ID}

	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{"date", dateValue}, {"id", m.Sender.ID}}
	_ = collection.FindOne(context.TODO(), filter).Decode(&find)

	// Если не находим, то создаём запись в БД
	if find.Weight == 0 {
		_, err := collection.InsertOne(context.TODO(), postValue)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Значение добавлено")

		// Иначе обновляем значение
	} else {
		update := bson.D{
			{"$set", bson.D{
				{"weight", weightValue},
			}},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Значение обновлено")
	}
}

func postOtherDate(m *tb.Message, collection *mongo.Collection, b *tb.Bot, date []string) {
	weightValue, err := strconv.ParseFloat(m.Text, 64)
	if err != nil {
		b.Send(m.Sender, "Ошибка ввода! Используйте только цифры. Пример: 60.5")
		return
	}
	dateValue := date[1] + "-" + date[2] + "-" + date[3]
	postValue := Weight{Date: dateValue, Weight: weightValue, ID: m.Sender.ID}

	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{"date", dateValue}, {"id", m.Sender.ID}}
	_ = collection.FindOne(context.TODO(), filter).Decode(&find)

	// Если не находим, то создаём запись в БД
	if find.Weight == 0 {
		_, err := collection.InsertOne(context.TODO(), postValue)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Значение добавлено")

		// Иначе обновляем значение
	} else {
		update := bson.D{
			{"$set", bson.D{
				{"weight", weightValue},
			}},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Значение обновлено")
	}
	b.Send(m.Sender, "Готов к получению команд.", menu)
}

func getValue(m *tb.Message, collection *mongo.Collection, b *tb.Bot,date []string) {
	dateValue := date[1] + "-" + date[2] + "-" + date[3]
	var find Weight

	// Ищем совпадение на основе полей даты и ID отправителя
	filter := bson.D{{"date", dateValue}, {"id", m.Sender.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&find)

	if err != nil {
		fmt.Println("Нет таких данных")
		_, err := b.Send(m.Sender, "Данные за "+dateValue+" отсутствуют.")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
		_, err = b.Send(m.Sender, "Вес за "+find.Date+" -- "+fmt.Sprintf("%.1f", find.Weight)+"кг")
		if err != nil {
			log.Fatal(err)
		}
}
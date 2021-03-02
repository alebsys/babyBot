package telegram

import (
	"context"
	"fmt"
	"log"
	"os"
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

func getHandler(m *tb.Message, collection *mongo.Collection, b *tb.Bot) {
	valueToSlice := strings.Split(m.Text, " ")
	dateValue := valueToSlice[0]

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
	} else {

		_, err := b.Send(m.Sender, "Вес за "+find.Date+" -- "+fmt.Sprintf("%.1f", find.Weight)+"кг")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func postHandler(m *tb.Message, collection *mongo.Collection, dateValue string, weightValue float64) {
	//valueToSlice := strings.Split(m.Text, " ")
	//dateValue := valueToSlice[1]
	//weightValue, _ := strconv.ParseFloat(valueToSlice[2], 64)

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

func graphHandler(m *tb.Message, collection *mongo.Collection, b *tb.Bot) {
	filter := bson.D{{"id", m.Sender.ID}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("collection.Find ERROR:", err)
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
	f, _ := os.Create("output.png")
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		log.Fatal("graph.Render ERROR:", err)
	}

	p := &tb.Photo{File: tb.FromDisk("output.png")}
	_, err = b.SendAlbum(m.Sender, tb.Album{p})
	if err != nil {
		log.Fatal("SendPhoto ERROR:", err)
	}
}

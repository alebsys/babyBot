package telegram

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	//"github.com/go-echarts/go-echarts/v2/opts"

	//"io"
	"log"
	//"os"
	"strings"
	//"github.com/go-echarts/go-echarts/v2/charts"
	//"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/wcharczuk/go-chart/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver@latest"

	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	w Weight
)

// getMenu ...
func getMenu(m *tb.Message) {
	_, _ = B.Send(m.Sender, "Что вы хотите посмотреть?", get)
	B.Handle(tb.OnText, func(m *tb.Message) {
		_, _ = B.Send(m.Sender, "Выберите один из пунктов меню.", get)
	})
}

// getMenuDate ...
func getMenuDate(m *tb.Message) {
	_, _ = B.Send(
		m.Sender,
		"Введите интересующую вас дату как `число/месяц/год`.\n\n"+
			"Пример: `21/10/20`.",
		&tb.SendOptions{
			ParseMode: tb.ModeMarkdown,
		},
		back)
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
		_, _ = b.Send(
			m.Sender,
			"*Неверный формат даты!*\n\nДата должна быть:\n\n"+
				"1. в формате `число/месяц/год`\n"+
				"2. за сегодняшнее или предыдущие числа\n\nПример `21/10/20`",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
		return errors.New("error from generateDate")
	}

	//_, _ = B.Send(m.Chat, rspMessage, &tb.SendOptions{
	//	DisableWebPagePreview: true,
	//	ParseMode:             tb.ModeMarkdown,
	//})
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
		_, err := b.Send(m.Sender, fmt.Sprintf("Данные за %v отсутствуют.", weight.Date))
		if err != nil {
			log.Printf("error from b.Send if collection.FindOne err != nil : %v", err)
		}

		return errors.New("error from getDate")
	}
	_, err = b.Send(m.Sender, fmt.Sprintf("Вес за %v -- %.1f кг", find.Date, find.Weight))
	if err != nil {
		log.Printf("error from b.Send if collection.FindOne err == nil : %v", err)
	}
	return nil
}

func getMenuGraph(m *tb.Message) {
	filter := bson.D{{Key: "id", Value: m.Sender.ID}}
	_, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("c.Find ERROR:", err)
	}

	values, max, _ := getGraphValues(w)

	graph := chart.BarChart{
		Title: "Динамика вашего веса",
		Background: chart.Style{
			Padding: chart.Box{
				Top:    50,
				Left:   25,
				Right:  25,
				Bottom: 50,
			},
		},
		YAxis: chart.YAxis{ // TODO добавить вывод значений по оси Y
			Range: &chart.ContinuousRange{
				Min: 0.0,
				Max: max + 30,
			},
		},
		Height:     512,
		BarWidth:   60,
		Bars:       values,
		BarSpacing: 500,
	}

	buffer := bytes.NewBuffer([]byte{})

	err = graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Printf("graph.Render ERROR: %v", err)
		_, _ = B.Send(
			m.Sender,
			"Мало данных для отображения графика (минимум 2 записи).\n\n*Добавьте еще!*",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
		return
	}

	p := &tb.Photo{File: tb.FromReader(buffer)}
	_, err = B.SendAlbum(m.Sender, tb.Album{p})
	if err != nil {
		log.Fatal("SendPhoto ERROR:", err)
	}

}

// getGraphValues ...
func getGraphValues(w Weight) ([]chart.Value, float64, error) {
	var v []chart.Value

	opt := options.Find()
	opt.SetSort(bson.D{{Key: "date", Value: 1}})
	sortCursor, err := collection.Find(context.TODO(), bson.D{{Key: "weight", Value: bson.D{{Key: "$gt", Value: 0}}}}, opt)
	if err != nil {
		log.Fatal(err)
	}

	var max float64
	for sortCursor.Next(context.TODO()) {
		// Decode the document
		if err := sortCursor.Decode(&w); err != nil {
			log.Fatal("cursor.Decode ERROR: ", err)
			return nil, 0, err
		}

		v = append(v,
			chart.Value{
				Label: w.Date,
				Value: w.Weight,
				Style: chart.Style{Hidden: false},
			})

		if w.Weight > max {
			max = w.Weight
		}
	}
	return v, max, nil
}

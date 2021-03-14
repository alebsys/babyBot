package telegram

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	l "github.com/alebsys/baby-bot/pkg/logs"
	"go.uber.org/zap"
	"strings"
	"github.com/wcharczuk/go-chart/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	getGraphError = "Мало данных для отображения графика (минимум 2 записи).\n\n*Добавьте еще!*"
	getMenuDateMess = "Введите интересующую вас дату как `число/месяц/год`.\n\nПример: `21/10/20`."
	generateDateWrong = "*Неверный формат даты!*\n\nДата должна быть:\n\n"+
		"1. в формате `число/месяц/год`\n"+
		"2. за сегодняшнее или предыдущие числа\n\nПример `21/10/20`"
)

var (
	w Weight
)

// getMenu ...
func getMenu(m *tb.Message) {
	if _, err := B.Send(m.Sender, "Что вы хотите посмотреть?", get); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}
	B.Handle(tb.OnText, func(m *tb.Message) {
		if _, err := B.Send(m.Sender, "Выберите один из пунктов меню.", get); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
	})
}

// getMenuDate ...
func getMenuDate(m *tb.Message) {
	if _, err := B.Send(m.Sender,getMenuDateMess, markdownOn,back); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}
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
		if _, err := b.Send(m.Sender,generateDateWrong, markdownOn); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
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

	if err := collection.FindOne(context.TODO(), filter).Decode(&find); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Info(err)


		if _, err := b.Send(m.Sender, fmt.Sprintf("Данные за %v отсутствуют.", weight.Date)); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}

		return errors.New("error from getDate")
	}
	if _, err := b.Send(m.Sender, fmt.Sprintf("Вес за %v -- %.1f кг", find.Date, find.Weight)); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}
	return nil
}

func getMenuGraph(m *tb.Message) {

	filter := bson.D{{Key: "id", Value: m.Sender.ID}}
	if _, err := collection.Find(context.TODO(), filter); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}

	values, max, _ := getGraphValues(w, m)

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

	if err := graph.Render(chart.PNG, buffer); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		if _, err = B.Send(m.Sender, getGraphError, markdownOn); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
		}
		return
	}

	p := &tb.Photo{File: tb.FromReader(buffer)}
	if _, err := B.SendAlbum(m.Sender, tb.Album{p}); err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}

}

// getGraphValues ...
func getGraphValues(w Weight, m *tb.Message) ([]chart.Value, float64, error) {
	var v []chart.Value

	opt := options.Find()
	opt.SetSort(bson.D{{Key: "date", Value: 1}})
	sortCursor, err := collection.Find(context.TODO(), bson.D{{Key: "weight", Value: bson.D{{Key: "$gt", Value: 0}}}}, opt)
	if err != nil {
		l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
	}

	var max float64
	for sortCursor.Next(context.TODO()) {
		// Decode the document
		if err := sortCursor.Decode(&w); err != nil {
			l.Sugar.With(zap.Int("clientID", m.Sender.ID), zap.Int("messageID", m.ID)).Error(err)
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

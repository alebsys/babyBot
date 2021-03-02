package telegram

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	//"time"
	"github.com/alebsys/baby-bot/config"
	"github.com/alebsys/baby-bot/pkg/db"

	//"github.com/spf13/viper"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Weight ...
type Weight struct {
	Date   string
	Weight float64
	ID     int
}

var (
	// Universal markup builders.
	//r.menu
	date []string
	menu = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	post = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	ages  = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	month = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	days  = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	back = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	get  = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

	// Reply buttons.
	btnPostMenu      = menu.Text("Ввести свой вес")
	btnPostToday     = post.Text("Да")
	btnPostOtherDate = post.Text("Нет")

	btnPost2020 = post.Text("2020")
	btnPost2021 = post.Text("2021")
	btnPost2022 = post.Text("2022")

	btnPostYan = post.Text("Январь")
	btnPostFeb = post.Text("Февраль")
	btnPostMar = post.Text("Март")
	btnPostApr = post.Text("Апрель")
	btnPostMay = post.Text("Май")
	btnPostJun = post.Text("Июнь")
	btnPostJul = post.Text("Июль")
	btnPostAug = post.Text("Август")
	btnPostSep = post.Text("Сентябрь")
	btnPostOct = post.Text("Октябрь")
	btnPostNov = post.Text("Ноябрь")
	btnPostDec = post.Text("Декабрь")

	btn01 = days.Text("01")
	btn02 = days.Text("02")
	btn03 = days.Text("03")
	btn04 = days.Text("04")
	btn05 = days.Text("05")
	btn06 = days.Text("06")
	btn07 = days.Text("07")
	btn08 = days.Text("08")
	btn09 = days.Text("09")
	btn10 = days.Text("10")
	btn11 = days.Text("11")
	btn12 = days.Text("12")
	btn13 = days.Text("13")
	btn14 = days.Text("14")
	btn15 = days.Text("15")
	btn16 = days.Text("16")
	btn17 = days.Text("17")
	btn18 = days.Text("18")
	btn19 = days.Text("19")
	btn20 = days.Text("20")
	btn21 = days.Text("21")
	btn22 = days.Text("22")
	btn23 = days.Text("23")
	btn24 = days.Text("24")
	btn25 = days.Text("25")
	btn26 = days.Text("26")
	btn27 = days.Text("27")
	btn28 = days.Text("28")
	btn29 = days.Text("29")
	btn30 = days.Text("30")
	btn31 = days.Text("31")

	btnGetMenu  = menu.Text("Посмотреть статистику")
	btnGetGraph = get.Text("Построить график")
	btnGetDate  = get.Text("Получить значение на Х дату")

	btnBackMenu = post.Text("Вернуться в меню")
)

// NewBot ...
func NewBot(configBot tb.Settings) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Println("Hello, I am babyBot!")

	collection := db.InitCollection()
	b, err := tb.NewBot(configBot)
	if err != nil {
		log.Fatal(err)
	}

	//b.Handle("/hello", func(m *tb.Message) {
	//	b.Send(m.Sender, "Hello World!")
	//})

	//b.Handle("/get", func(m *tb.Message) {
	//	getHandler(m, collection, b)
	//})

	//b.Handle("/delete", func(m *tb.Message) {
	//	deleteHandler(m, collection, b)
	//})

	//b.Handle("/graph", func(m *tb.Message) {
	//	graphHandler(m, collection, b)
	//})

	//b.Handle("/post", func(m *tb.Message) {
	//	valueToSlice := strings.Split(m.Text, " ")
	//	dateValue := valueToSlice[1]
	//	weightValue, _ := strconv.ParseFloat(valueToSlice[2], 64)
	//	postHandler(m, collection,dateValue,weightValue)
	//})

	//b.Handle("/graph", func(m *tb.Message) {
	//	graphHandler(m, collection, b)
	//})
	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Привет!", menu)
	})

	menu.Reply(
		menu.Row(btnPostMenu),
		menu.Row(btnGetMenu),
	)

	post.Reply(
		post.Row(btnPostToday),
		post.Row(btnPostOtherDate),
		post.Row(btnBackMenu),
	)
	ages.Reply(
		ages.Row(btnPost2020),
		ages.Row(btnPost2021),
		ages.Row(btnPost2022),
		ages.Row(btnBackMenu),
	)
	month.Reply(
		post.Row(btnPostYan),
		post.Row(btnPostFeb),
		post.Row(btnPostMar),
		post.Row(btnPostApr),
		post.Row(btnPostMay),
		post.Row(btnPostJun),
		post.Row(btnPostJul),
		post.Row(btnPostAug),
		post.Row(btnPostSep),
		post.Row(btnPostOct),
		post.Row(btnPostNov),
		post.Row(btnPostDec),
		post.Row(btnBackMenu),
	)

	days.Reply(
		post.Row(btn01),
		post.Row(btn02),
		post.Row(btn03),
		post.Row(btn04),
		post.Row(btn05),
		post.Row(btn06),
		post.Row(btn07),
		post.Row(btn08),
		post.Row(btn09),
		post.Row(btn10),
		post.Row(btn11),
		post.Row(btn12),
		post.Row(btn13),
		post.Row(btn14),
		post.Row(btn15),
		post.Row(btn16),
		post.Row(btn17),
		post.Row(btn18),
		post.Row(btn19),
		post.Row(btn20),
		post.Row(btn21),
		post.Row(btn22),
		post.Row(btn23),
		post.Row(btn24),
		post.Row(btn25),
		post.Row(btn26),
		post.Row(btn27),
		post.Row(btn28),
		post.Row(btn29),
		post.Row(btn30),
		post.Row(btn31),
		post.Row(btnBackMenu),
	)

	back.Reply(
		back.Row(btnBackMenu),
	)

	get.Reply(
		get.Row(btnGetGraph),
		get.Row(btnGetDate),
		get.Row(btnBackMenu),
	)

	// Кнопки для добавления данных
	b.Handle(&btnPostMenu, func(m *tb.Message) {
		b.Send(m.Sender, "Вводим данные за сегодня?", post)
	})

	b.Handle(&btnPostToday, func(m *tb.Message) {
		b.Send(m.Sender, "Введите свой весь в кг. Пример: 60.5", back)

		b.Handle(tb.OnText, func(m *tb.Message) {
			//now := time.Now()
			layout := "2006-01-02"
			valueToSlice := strings.Split(m.Text, " ")
			dateValue := time.Now().Format(layout)
			weightValue, _ := strconv.ParseFloat(valueToSlice[0], 64)
			postHandler(m, collection, dateValue, weightValue)

			b.Send(m.Sender, "Ваш сегодняшний вес: "+valueToSlice[0]+" кг.", menu)
		})

	})

	b.Handle(&btnPostOtherDate, func(m *tb.Message) {
		//b.Send(m.Sender, "Введи дату замера (год-месяц-день) и свой вес в кг. Пример: 2020-12-31 65", back)
		b.Send(m.Sender, "Выберите год.", ages)

	})

	// Кнопки для получения данных
	b.Handle(&btnGetMenu, func(m *tb.Message) {
		b.Send(m.Sender, "Что вы хотите посмотреть?", get)
	})

	b.Handle(&btnGetGraph, func(m *tb.Message) {
		b.Send(m.Sender, "", get)
		graphHandler(m, collection, b)
	})

	b.Handle(&btnGetDate, func(m *tb.Message) {
		b.Send(m.Sender, "Введи дату замера (год-месяц-день). Пример: 2020-12-31", get)
		b.Handle(tb.OnText, func(m *tb.Message) {
			getHandler(m, collection, b)
		})
	})

	// Кнопки для возврата в меню
	b.Handle(&btnBackMenu, func(m *tb.Message) {
		b.Send(m.Sender, "Возвращаемся в меню", menu)
	})

	// Выбираем год
	b.Handle(&btnPost2020, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите месяц.", month)
	})

	b.Handle(&btnPost2021, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите месяц.", month)
	})

	b.Handle(&btnPost2022, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите месяц.", month)
	})

	// Выбираем месяц
	b.Handle(&btnPostYan, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostFeb, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostMar, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})
	b.Handle(&btnPostApr, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostMay, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostJun, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostJul, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostAug, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostSep, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostOct, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostNov, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	b.Handle(&btnPostDec, func(m *tb.Message) {
		date = append(date, m.Text)
		b.Send(m.Sender, "Выберите число.", days)
	})

	// // // /// // /// //// Числа
	//b.Send(m.Sender, "Данные за "+date[0]+"-"+date[1]+"-"+date[2], back)

	b.Handle(&btn01, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})

	b.Handle(&btn02, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn03, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn04, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn05, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn06, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn07, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn08, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn09, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn10, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn11, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn12, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn13, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn14, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn15, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn16, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn17, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn18, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn19, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn20, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn21, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn22, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn23, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn24, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn25, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn26, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn27, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn28, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn29, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn30, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})
	b.Handle(&btn31, func(m *tb.Message) {
		date = append(date, m.Text)

		b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			weightValue, _ := strconv.ParseFloat(m.Text, 64)
			dateValue := date[0] + "-" + date[1] + "-" + date[2]
			postHandler(m, collection, dateValue, weightValue)
		})
	})

	b.Start()
}

package telegram

import (
	"fmt"
	"log"
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
	dateLayout = "2006-Jan-02"
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

	btnPostYan = post.Text("Yan")
	btnPostFeb = post.Text("Feb")
	btnPostMar = post.Text("Mar")
	btnPostApr = post.Text("Apr")
	btnPostMay = post.Text("May")
	btnPostJun = post.Text("Jun")
	btnPostJul = post.Text("Jul")
	btnPostAug = post.Text("Aug")
	btnPostSep = post.Text("Sep")
	btnPostOct = post.Text("Oct")
	btnPostNov = post.Text("Nov")
	btnPostDec = post.Text("Dec")

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

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Привет!", menu)
	})

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

	// Главное меню
	menu.Reply(
		menu.Row(btnPostMenu),
		menu.Row(btnGetMenu),
	)
	// Добавить данные
	b.Handle(&btnPostMenu, func(m *tb.Message) {
		date = []string{}
		b.Send(m.Sender, "Вводим данные за сегодня?", post)
	})

	// Получить данные
	b.Handle(&btnGetMenu, func(m *tb.Message) {
		date = []string{}
		b.Send(m.Sender, "Что вы хотите посмотреть?", get)
	})

	//b.Handle(&btnPostToday, func(m *tb.Message) {
	//	b.Send(m.Sender, "Введите свой вес в кг. Пример: 60.5", back)
	//	b.Handle(tb.OnText, func(m *tb.Message) {
	//
	//		valueToSlice := strings.Split(m.Text, " ")
	//		weightValue, err := strconv.ParseFloat(valueToSlice[0], 64)
	//		if err != nil {
	//			b.Send(m.Sender, "Введите корректное значение веса. Пример: 60.5")
	//		}
	//		postHandler(m, collection, time.Now().Format(dateLayout), weightValue)
	//		b.Send(m.Sender, "Ваш сегодняшний вес: "+valueToSlice[0]+" кг.", menu)
	//	})
	//})

	// Добавить данные за сегодня
	b.Handle(&btnPostToday, func(m *tb.Message) {
		b.Send(m.Sender, "Введите свой вес в кг. Пример: 60.5", back)
		b.Handle(tb.OnText, func(m *tb.Message) {
			postToday(m, collection, b, time.Now().Format(dateLayout))
			b.Send(m.Sender, "Ваш сегодняшний вес записан.", menu)
		})
	})

	b.Handle(&btnPostOtherDate, func(m *tb.Message) {
		date = append(date, "postDate")
		b.Send(m.Sender, "Выберите год.", ages)

	})

	// Кнопки для получения данных
	b.Handle(&btnGetGraph, func(m *tb.Message) {
		b.Send(m.Sender, "", get)
		graphHandler(m, collection, b)
	})

	b.Handle(&btnGetDate, func(m *tb.Message) {
		date = append(date, "getDate")
		b.Send(m.Sender, "Выберите год", ages)
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

	// Кнопки чисел
	//b.Handle(&btn01, func(m *tb.Message) {
	//	date = append(date, m.Text)
	//	if date[0] == "postDate" {
	//		postDate(m, collection, b, date)
	//	}
	//	getHandler(m, collection, b, date)
	//	b.Send(m.Sender, "Идем обратно в меню.", menu)
	//
	//})
	b.Handle(&btn01, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})


	b.Handle(&btn02, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})
	
	b.Handle(&btn03, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn04, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn05, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn06, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn07, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn08, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn09, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn10, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn11, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn12, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn13, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn14, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn15, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	
	})

	b.Handle(&btn16, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn17, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn18, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn19, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn20, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn21, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn22, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn23, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn24, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn25, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn26, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn27, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn28, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn29, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn30, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})

	b.Handle(&btn31, func(m *tb.Message) {
		date = append(date, m.Text)
		if date[0] == "postDate" {
			b.Send(m.Sender, "Введите ваш вес в кг. Пример: 60.5", back)
			b.Handle(tb.OnText, func(m *tb.Message) {
				postOtherDate(m, collection,b, date)
			})
			return
		}
		getValue(m, collection, b, date)
		b.Send(m.Sender, "Готов к получению команд.", menu)
	})


	b.Start()
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: time.Second * 10},
	})

	if err != nil {
		log.Fatalln(err)
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Привет, я всего лишь бот для расчета стоимости поездки)")
	})

	b.Handle("/cost", func(m *tb.Message) {
		data := strings.Split(m.Text[6:], " ")
		if len(data) < 3 {
			b.Send(m.Sender, "Дружок, а ты походу туговат, правила ввода описаны в /start")
			return
		}
		distance, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			b.Send(m.Sender, "Неправильный ввод!")
			return
		}
		consumption, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			b.Send(m.Sender, "Неправильный ввод!")
			return
		}
		costPerLiter, err := strconv.ParseFloat(data[2], 64)
		if err != nil {
			b.Send(m.Sender, "Неправильный ввод!")
			return
		}
		spent := distance * consumption / 100
		cost := spent * costPerLiter
		text := fmt.Sprintf(
			"Кол-во потраченого топлива: <b>%.2f л.</b>\nКол-во денег: <b>%.2f грн.</b>", spent, cost,
		)
		b.Send(m.Sender, text, &tb.SendOptions{ParseMode: "html"})
	})

	b.Start()
}

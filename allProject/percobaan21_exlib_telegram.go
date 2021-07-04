package main

import (
	"time"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
	"fmt"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "528831631:AAFj16ScMOYagEvD3uatkrtku89FnD8govs",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world, psmk golang")
	})

	b.Handle("/reg", func(m *tb.Message) {
		fmt.Println(m.Text)
		b.Send(m.Sender, "Anda berhasil melakukan registrasi!")
	})

	b.Start()
}
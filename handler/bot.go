package handler

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func (h *BotHandler) StartBot(apiToken string) {
	pref := tele.Settings{
		Token:  apiToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Привет!\nДавай сыграем в города! Назови город в России")
	})

	bot.Handle(tele.OnText, h.PlayGame)

	log.Print("listen to telegram api")
	bot.Start()
}

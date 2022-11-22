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

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	hintBtn := menu.Text("Подсказка")
	menu.Reply(menu.Row(hintBtn))

	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Привет!\nДавай сыграем в города! Назови город в России", menu)
	})

	bot.Handle("/restart", h.RestartGame)
	bot.Handle(tele.OnText, h.PlayGame)
	bot.Handle(&hintBtn, h.GetHint)

	log.Print("listen to telegram api")
	bot.Start()
}

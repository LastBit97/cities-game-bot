package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/LastBit97/cities-game-bot/service"
	tele "gopkg.in/telebot.v3"
)

var lastCities map[int64]string

type BotHandler struct {
	cityService service.CitiesGame
}

func NewBotHandler(service service.CitiesGame) BotHandler {
	return BotHandler{service}
}

func (h *BotHandler) PlayGame(ctx tele.Context) error {
	city := ctx.Text()
	chatId := ctx.Chat().ID

	if lastCities == nil {
		lastCities = make(map[int64]string)
	}

	var lastCity string
	if val, ok := lastCities[chatId]; ok {
		lastCity = val
	}

	if !h.cityService.Exists(city) {
		if err := ctx.Send(cityNotExist); err != nil {
			return err
		}
		return nil
	}

	log.Println(lastCity)
	log.Println(city)

	if lastCity != "" {
		if !h.cityService.CheckCity(lastCity, city) {
			letter := h.cityService.GetLastChar(lastCity)
			msg := fmt.Sprintf(letterResponse, letter)
			if err := ctx.Send(msg); err != nil {
				return err
			}
			return nil
		}
	}

	if !h.cityService.Contains(city) {
		if err := ctx.Send(cityBeenUse); err != nil {
			return err
		}
		return nil
	}

	cityReply, err := h.cityService.GetRandomCity(city)
	if err != nil {
		if err := ctx.Send(citiesEnded); err != nil {
			return err
		}
		return err
	}

	letter := h.cityService.GetLastChar(cityReply)
	letterMsg := fmt.Sprintf(letterResponse, letter)

	durationReply := time.Duration(2) * time.Second
	time.Sleep(durationReply)

	if err := ctx.Send("Думаю..."); err != nil {
		return err
	}

	time.Sleep(durationReply)

	if err := ctx.Send("Очень сильно думаю..."); err != nil {
		return err
	}

	time.Sleep(durationReply)

	if err := ctx.Send(cityReply); err != nil {
		return err
	}

	if err := ctx.Send(letterMsg); err != nil {
		return err
	}

	lastCities[chatId] = cityReply
	h.cityService.DeleteCity(city)
	h.cityService.DeleteCity(cityReply)

	return nil
}

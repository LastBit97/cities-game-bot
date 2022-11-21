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
	log.Printf("player named city: %s", city)

	if !h.cityService.CheckList(chatId) {
		h.cityService.NewList(chatId)
	}

	if lastCities == nil {
		lastCities = make(map[int64]string)
	}

	check, err := h.checkPlayerMsg(chatId, city, ctx)
	if !check {
		return err
	}

	h.cityService.DeleteCity(city, chatId)

	cityReply, err := h.cityService.GetRandomCity(city, chatId)
	if err != nil {
		if err := ctx.Send(citiesEnded); err != nil {
			return err
		}
		return err
	}
	durationReply := time.Second
	time.Sleep(durationReply)

	if err := ctx.Send(cityReply); err != nil {
		return err
	}
	log.Printf("bot reply: %s", cityReply)

	letter := h.cityService.GetLastChar(cityReply)
	letterMsg := fmt.Sprintf(letterResponse, letter)
	if err := ctx.Send(letterMsg); err != nil {
		return err
	}

	lastCities[chatId] = cityReply
	h.cityService.DeleteCity(cityReply, chatId)
	return nil
}

func (h *BotHandler) checkPlayerMsg(chatId int64, city string, ctx tele.Context) (bool, error) {
	var lastCity string
	if val, ok := lastCities[chatId]; ok {
		lastCity = val
	}

	if !h.cityService.Exists(city) {
		if err := ctx.Send(cityNotExist); err != nil {
			return false, err
		}
		return false, nil
	}

	if lastCity != "" {
		if !h.cityService.CheckCity(lastCity, city) {
			log.Println("letters don't match")
			letter := h.cityService.GetLastChar(lastCity)
			msg := fmt.Sprintf(letterResponse, letter)
			if err := ctx.Send(msg); err != nil {
				return false, err
			}
			return false, nil
		}
	}

	if !h.cityService.Contains(city, chatId) {
		if err := ctx.Send(cityBeenUse); err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

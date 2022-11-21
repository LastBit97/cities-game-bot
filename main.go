package main

import (
	"log"

	"github.com/LastBit97/cities-game-bot/configs"
	"github.com/LastBit97/cities-game-bot/handler"
	"github.com/LastBit97/cities-game-bot/service"
)

var telegramBot handler.BotHandler
var cityService service.CitiesGame

func init() {
	cityService = service.NewCitiesGame()
	telegramBot = handler.NewBotHandler(cityService)
}

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	telegramBot.StartBot(config.ApiToken)
}

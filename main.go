package main

import (
	"log"

	"github.com/LastBit97/cities-game-bot/configs"
	"github.com/LastBit97/cities-game-bot/handler"
	"github.com/LastBit97/cities-game-bot/model"
	"github.com/LastBit97/cities-game-bot/service"
	"github.com/LastBit97/cities-game-bot/utils"
)

var telegramBot handler.BotHandler
var cityService service.CitiesGame

func main() {
	var cities []model.City

	if err := utils.ReadAndUnmarshal("russian-cities.json", &cities); err != nil {
		log.Println(err)
	}

	var citiesName []string
	for _, city := range cities {
		citiesName = append(citiesName, city.Name)
	}

	cityService = service.NewCitiesGame(citiesName)
	telegramBot = handler.NewBotHandler(cityService)

	// citiesGame := service.NewCitiesGame()
	// fmt.Println(citiesGame.GetCities())
	// citiesGame.DeleteCity("Москва")
	// fmt.Println(citiesGame.GetCities())
	// fmt.Println(citiesGame.Contains("Архангельск"))
	// result := citiesGame.GetRandomCity("Москва")
	// fmt.Println(result)

	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	telegramBot.StartBot(config.ApiToken)

}

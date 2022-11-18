package main

import (
	"fmt"

	"github.com/LastBit97/cities-game-bot/service"
)

func main() {
	// var cities []model.City

	// if err := utils.ReadAndUnmarshal("russian-cities.json", &cities); err != nil {
	// 	log.Println(err)
	// }

	// var citiesName []string
	// for _, city := range cities {
	// 	citiesName = append(citiesName, city.Name)
	// }

	citiesGame := service.NewCitiesGame()
	fmt.Println(citiesGame.GetCities())
	citiesGame.DeleteCity("Москва")
	fmt.Println(citiesGame.GetCities())
	fmt.Println(citiesGame.Contains("Архангельск"))
	result := citiesGame.GetRandomCity("Москва")
	fmt.Println(result)

}

package service

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/LastBit97/cities-game-bot/model"
	"github.com/LastBit97/cities-game-bot/utils"
)

type citiesGameImpl struct {
	cities []string
}

func NewCitiesGame() CitiesGame {
	cities := []string{"Москва", "Архангельск", "Кирсанов", "Владивосток", "Кострома", "Астрахань"}
	return &citiesGameImpl{cities}
}

func (cg *citiesGameImpl) DeleteCity(cityName string) {
	for i, city := range cg.cities {
		if city == cityName {
			cg.cities = append(cg.cities[:i], cg.cities[i+1:]...)
		}
	}
}

func (cg *citiesGameImpl) Exists(cityName string) bool {
	var cities []model.City

	if err := utils.ReadAndUnmarshal("russian-cities.json", &cities); err != nil {
		log.Println(err)
	}

	var fullListCities []string
	for _, city := range cities {
		fullListCities = append(fullListCities, city.Name)
	}

	for _, city := range fullListCities {
		if city == cityName {
			return true
		}
	}
	return false

}

func (cg *citiesGameImpl) Contains(cityName string) bool {
	for _, city := range cg.cities {
		if city == cityName {
			return true
		}
	}
	return false

}

func (cg *citiesGameImpl) GetRandomCity(cityName string) string {
	cities := cg.getCorrectCities(cityName)
	rand.Seed(time.Now().Unix())
	return cities[rand.Intn(len(cities))]
}

func (cg *citiesGameImpl) getCorrectCities(cityName string) []string {
	cityByRune := []rune(cityName)
	lastChar := string(cityByRune[len(cityByRune)-1:])

	var correctCities []string
	for _, city := range cg.cities {
		cityToRunes := []rune(city)
		firstChar := string(cityToRunes[0:1])
		if strings.EqualFold(firstChar, lastChar) {
			correctCities = append(correctCities, city)
		}
	}
	return correctCities
}

func (cg *citiesGameImpl) GetCities() []string {
	return cg.cities
}

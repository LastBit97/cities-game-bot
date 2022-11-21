package service

import (
	"errors"
	"fmt"
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

func NewCitiesGame(cities []string) CitiesGame {
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

func (cg *citiesGameImpl) GetRandomCity(cityName string) (string, error) {
	cities, err := cg.getCorrectCities(cityName)
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().Unix())
	return cities[rand.Intn(len(cities))], nil
}

func (cg *citiesGameImpl) getCorrectCities(cityName string) ([]string, error) {
	lastChar := cg.GetLastChar(cityName)

	var correctCities []string
	for _, city := range cg.cities {
		firstChar := cg.getFirstChar(city)
		if strings.EqualFold(firstChar, lastChar) {
			correctCities = append(correctCities, city)
		}
	}
	if correctCities == nil {
		return nil, errors.New("no cities found")
	}
	return correctCities, nil
}

func (cg *citiesGameImpl) getFirstChar(city string) string {
	cityToRunes := []rune(city)
	firstChar := string(cityToRunes[0:1])
	return firstChar
}

func (cg *citiesGameImpl) GetLastChar(cityName string) string {
	if cityName == "" {
		return ""
	}
	cityByRune := []rune(cityName)
	lastChar := string(cityByRune[len(cityByRune)-1:])
	if lastChar == "ь" || lastChar == "ы" {
		cityName = string(cityByRune[:len(cityByRune)-1])
		fmt.Println(cityName)
		return cg.GetLastChar(cityName)
	}
	return lastChar
}

func (cg *citiesGameImpl) CheckCity(lastCity string, currentCity string) bool {
	lastChar := cg.GetLastChar(lastCity)
	firstChar := cg.getFirstChar(currentCity)
	return strings.EqualFold(lastChar, firstChar)
}

func (cg *citiesGameImpl) GetCities() []string {
	return cg.cities
}

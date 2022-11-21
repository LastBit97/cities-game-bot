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
	fullListCities []string
	cities         []string
}

func NewCitiesGame() CitiesGame {
	fullListCities := getFullListCities()
	cities := getFullListCities()
	return &citiesGameImpl{fullListCities, cities}
}

func (cg *citiesGameImpl) GetCities() []string {
	return cg.cities
}

func (cg *citiesGameImpl) DeleteCity(cityName string) {
	for i, city := range cg.cities {
		if city == cityName {
			cg.cities = append(cg.cities[:i], cg.cities[i+1:]...)
		}
	}
	log.Printf("remove city: %s from list", cityName)
}

func (cg *citiesGameImpl) Exists(cityName string) bool {
	for _, city := range cg.fullListCities {
		if city == cityName {
			return true
		}
	}
	log.Printf("city: %s doesn't exist", cityName)
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

func (cg *citiesGameImpl) CheckCity(lastCity string, currentCity string) bool {
	lastChar := cg.GetLastChar(lastCity)
	firstChar := cg.getFirstChar(currentCity)
	return strings.EqualFold(lastChar, firstChar)
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

func (cg *citiesGameImpl) getFirstChar(city string) string {
	if city == "" {
		return ""
	}
	cityToRunes := []rune(city)
	firstChar := string(cityToRunes[0:1])
	return firstChar
}

func getFullListCities() []string {
	var cities []model.City

	if err := utils.ReadAndUnmarshal("russian-cities.json", &cities); err != nil {
		log.Fatalln(err)
	}

	var russianCities []string
	for _, city := range cities {
		russianCities = append(russianCities, city.Name)
	}
	return russianCities
}

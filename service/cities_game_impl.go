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
	citiesMap      map[int64][]string
}

func NewCitiesGame() CitiesGame {
	fullListCities := getListCities()
	citiesMap := make(map[int64][]string)
	return &citiesGameImpl{fullListCities, citiesMap}
}

func (cg *citiesGameImpl) CheckList(chatId int64) bool {
	if _, ok := cg.citiesMap[chatId]; ok {
		return true
	}
	return false
}

func (cg *citiesGameImpl) NewList(chatId int64) {
	cities := getListCities()
	cg.citiesMap[chatId] = cities
}

func (cg *citiesGameImpl) GetCities(chatId int64) []string {
	return cg.citiesMap[chatId]
}

func (cg *citiesGameImpl) DeleteCity(cityName string, chatId int64) {
	cities := cg.citiesMap[chatId]

	for i, city := range cities {
		if city == cityName {
			cities = append(cities[:i], cities[i+1:]...)
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

func (cg *citiesGameImpl) Contains(cityName string, chatId int64) bool {
	cities := cg.citiesMap[chatId]
	for _, city := range cities {
		if city == cityName {
			return true
		}
	}
	return false
}

func (cg *citiesGameImpl) GetRandomCity(cityName string, chatId int64) (string, error) {
	cities, err := cg.getCorrectCities(cityName, chatId)
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().Unix())
	return cities[rand.Intn(len(cities))], nil
}

func (cg *citiesGameImpl) getCorrectCities(cityName string, chatId int64) ([]string, error) {
	lastChar := cg.GetLastChar(cityName)
	cities := cg.citiesMap[chatId]

	var correctCities []string
	for _, city := range cities {
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

func getListCities() []string {
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

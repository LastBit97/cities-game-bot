package service

type CitiesGame interface {
	DeleteCity(cityName string, chatId int64)
	Exists(cityName string) bool
	Contains(cityName string, chatId int64) bool
	GetRandomCity(cityName string, chatId int64) (string, error)
	GetCities(chatId int64) []string
	CheckCity(lastCity string, currentCity string) bool
	GetLastChar(cityName string) string
	NewList(chatId int64)
	CheckList(chatId int64) bool
}

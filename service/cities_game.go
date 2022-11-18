package service

type CitiesGame interface {
	DeleteCity(cityName string)
	Exists(cityName string) bool
	Contains(cityName string) bool
	GetRandomCity(cityName string) string
	GetCities() []string
	CheckCity(lastCity string, currentCity string) bool
}

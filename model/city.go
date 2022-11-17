package model

type City struct {
	Coords     Coordinates `json:"coords"`
	District   string      `json:"district"`
	Name       string      `json:"name"`
	Population uint        `json:"population"`
	Subject    string      `json:"subject"`
}

type Coordinates struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}

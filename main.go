package main

import (
	"fmt"
	"log"

	"github.com/LastBit97/cities-game-bot/model"
	"github.com/LastBit97/cities-game-bot/utils"
)

func main() {

	var cities []model.City

	if err := utils.ReadAndUnmarshal("russian-citis.json", &cities); err != nil {
		log.Println(err)
	}

	for _, city := range cities {
		fmt.Print(city.Name, " ")
	}

}

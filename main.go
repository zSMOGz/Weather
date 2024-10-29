package main

import (
	"flag"
	"fmt"

	"main/geo"
	"main/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*format)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData.City)
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}

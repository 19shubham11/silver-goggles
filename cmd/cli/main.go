package main

import (
	"19shubham11/weather-cli/pkg/weather"
	"fmt"
)

func main() {
	res, _ := weather.GetCurrentWeather()
	fmt.Printf("%+v\n", res)

	fmt.Println(res.Weather[0].Conditions)
}

package main

import (
	config "19shubham11/weather-cli/config"
	weather "19shubham11/weather-cli/pkg/weather"
	"fmt"
)

func main() {
	appConfig := config.LoadAppConfig()
	weatherAPI := weather.WeatherAPI{
		Conf: appConfig,
	}

	res, _ := weatherAPI.GetCurrentWeather()
	fmt.Printf("%+v\n", res)

	fmt.Println(res.Weather[0].Conditions)
}

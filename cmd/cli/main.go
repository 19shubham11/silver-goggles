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

	res, err := weatherAPI.GetCurrentWeather("Berlin")
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
}

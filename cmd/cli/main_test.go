package main

import (
	config "19shubham11/weather-cli/config"
	weather "19shubham11/weather-cli/pkg/weather"
	"fmt"
	"io"
	"os"
	"testing"
)

var weatherAPI weather.WeatherAPI

func TestMain(m *testing.M) {

	appConfig := config.LoadAppConfig()
	weatherAPI = weather.WeatherAPI{
		Conf: appConfig,
	}
	code := m.Run()
	os.Exit(code)
}

func TestIntegration(t *testing.T) {
	args := []string{"current", "-city", "berlin"}
	os.Args = args
	var out io.Writer = os.Stdout

	err := setupCLI(args, weatherAPI, out)
	fmt.Println("err", err)
}

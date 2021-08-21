package main

import (
	config "19shubham11/weather-cli/config"
	weather "19shubham11/weather-cli/pkg/weather"
	"bytes"
	"os"
	"strings"
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
	out := bytes.NewBuffer(nil)

	err := setupCLI(args, weatherAPI, out)
	if err != nil {
		t.Errorf("Error getting weather!")
	}

	output := out.String()
	if !strings.Contains(output, "Current weather for") && !strings.Contains(output, "Feels like") && !strings.Contains(output, "expect") {
		t.Error("Formatting error!", output)
	}
}

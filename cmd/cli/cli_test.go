package main

import (
	weather "19shubham11/weather-cli/pkg/weather"
	"flag"
	"fmt"
	"os"
	"testing"
)

type mockWeather struct{}

func (m mockWeather) GetCurrentWeather(cityName string) (*weather.CurrentWeather, error) {
	return &weather.CurrentWeather{}, nil
}

var mockCommand WeatherCommand

const CommandName = "test"

func TestMain(m *testing.M) {
	mockCommand = WeatherCommand{
		fs:  flag.NewFlagSet(CommandName, flag.ExitOnError),
		api: &mockWeather{},
	}

	var _ = func() bool {
		testing.Init()
		return true
	}()
	code := m.Run()
	os.Exit(code)
}

func TestName(t *testing.T) {
	t.Run("should return the name of the command", func(t *testing.T) {
		name := mockCommand.Name()

		if name != CommandName {
			t.Errorf("expected %s, got %s", CommandName, name)
		}
	})
}

func TestInit(t *testing.T) {

	var tests = []struct {
		name string
		inp  []string
	}{
		{"should return proper error when flag is not present", []string{}},
		{"should return error if an unknown flag is set", []string{"cityName=london"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualErr := mockCommand.Init(tt.inp)
			fmt.Println("actualErr", actualErr)
			if actualErr == nil {
				t.Errorf("Expected err %v", actualErr)
			}
		})
	}
}

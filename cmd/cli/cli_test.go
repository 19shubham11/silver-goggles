package main

import (
	"bytes"
	"errors"
	"flag"
	"testing"

	"19shubham11/weather-cli/cmd/pkg/weather"
	"19shubham11/weather-cli/cmd/test/helpers"
)

type mockWeather struct {
	mockResp weather.CurrentWeather
	mockErr  error
}

func (m mockWeather) GetCurrentWeather(cityName string) (*weather.CurrentWeather, error) {
	return &m.mockResp, m.mockErr
}

var mockCommand WeatherCommand

const testCommand = "test"

var errWeatherMock = errors.New("weather error")

func init() {
	mockCommand = WeatherCommand{
		fs:     flag.NewFlagSet(testCommand, flag.ExitOnError),
		api:    &mockWeather{},
		output: bytes.NewBuffer(nil),
	}
}

func TestName(t *testing.T) {
	name := mockCommand.Name()

	if name != testCommand {
		t.Errorf("expected %s, got %s", testCommand, name)
	}
}

func TestInit(t *testing.T) {
	var tests = map[string]struct {
		inp []string
	}{
		"returns proper error when flag is not present": {[]string{}},
		"returns error if an unknown flag is set":       {[]string{"cityName=london"}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actualErr := mockCommand.Init(tt.inp)
			if actualErr == nil {
				t.Errorf("Expected err %v", actualErr)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := map[string]struct {
		command     WeatherCommand
		expectedOut []string
		expectedErr error
	}{
		"weather/current success": {
			WeatherCommand{
				fs: flag.NewFlagSet(CommandCurrentWeather, flag.ExitOnError),
				api: &mockWeather{
					mockResp: weather.CurrentWeather{
						Weather: weather.Description{
							{
								Summary:     "clear blue skies",
								Description: "zzz",
							},
						},
						Values: weather.Temperature{
							Temp:      22,
							FeelsLike: 22.5,
						},
					},
					mockErr: nil,
				},
				city: "Berlin",
			},
			[]string{"Current weather for", "Feels like", "Expect"},
			nil,
		},
		"weather/current error": {
			WeatherCommand{
				fs: flag.NewFlagSet(CommandCurrentWeather, flag.ExitOnError),
				api: &mockWeather{
					mockResp: weather.CurrentWeather{},
					mockErr:  errWeatherMock,
				},
				city: "Berlin",
			},
			[]string{""},
			errWeatherMock,
		},
		"weather/weekly": {
			WeatherCommand{
				fs:   flag.NewFlagSet(CommandWeeklyWeather, flag.ExitOnError),
				api:  &mockWeather{},
				city: "Berlin",
			},
			[]string{"Not implemented yet"},
			nil,
		},
		"help": {
			WeatherCommand{
				fs:   flag.NewFlagSet(CommandHelp, flag.ExitOnError),
				api:  &mockWeather{},
				city: "Berlin",
			},
			[]string{"$ current -city=Berlin", "weekly  -city=Toronto"},
			nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			tt.command.output = buffer

			err := tt.command.Run()
			stringOutput := buffer.String()

			helpers.AssertError(t, tt.expectedErr, err)
			helpers.AssertSubstrings(t, stringOutput, tt.expectedOut)
		})
	}
}

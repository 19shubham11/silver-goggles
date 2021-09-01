package main

import (
	"bytes"
	"os"
	"testing"

	"19shubham11/weather-cli/config"
	"19shubham11/weather-cli/pkg/weather"
	"19shubham11/weather-cli/test/helpers"
)

var weatherAPI weather.OpenWeatherAPI

func TestMain(m *testing.M) {
	appConfig := config.LoadAppConfig()
	weatherAPI = weather.OpenWeatherAPI{
		Conf: appConfig,
	}
	code := m.Run()
	os.Exit(code)
}

func TestIntegration(t *testing.T) {
	tests := map[string]struct {
		args            []string
		out             *bytes.Buffer
		expectedStrings []string
		expectedErr     error
	}{
		"current": {
			[]string{"current", "-city", "berlin"},
			bytes.NewBuffer(nil),
			[]string{"Current weather for", "Feels like", "Expect"},
			nil,
		},
		"weekly": {
			[]string{"weekly", "-city", "berlin"},
			bytes.NewBuffer(nil),
			[]string{"Not implemented yet"},
			nil,
		},

		"help": {
			[]string{"help"},
			bytes.NewBuffer(nil),
			[]string{"$ current -city=Berlin", "weekly  -city=Toronto"},
			nil,
		},

		"empty city flag": {
			[]string{"weekly", "-city", ""},
			bytes.NewBuffer(nil),
			[]string{""},
			ErrorCityMissing,
		},
		"unsupported command": {
			[]string{"hourly", "-city", "berlin"},
			bytes.NewBuffer(nil),
			[]string{""},
			ErrorUnknownCommand,
		},
		"insufficient args": {
			[]string{},
			bytes.NewBuffer(nil),
			[]string{""},
			ErrorInsufficientArgs,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			os.Args = tt.args
			err := setupCLI(tt.args, weatherAPI, tt.out)
			helpers.AssertError(t, tt.expectedErr, err)
			stringOutput := tt.out.String()
			helpers.AssertSubstrings(t, stringOutput, tt.expectedStrings)
		})
	}
}

package main

import (
	config "19shubham11/weather-cli/config"
	weather "19shubham11/weather-cli/pkg/weather"
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
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

func assertError(t *testing.T, expected, got error) {
	t.Helper()

	if !errors.Is(got, expected) {
		t.Fatalf("Expected error %v got %v", expected, got)
	}
}

func assertConsoleOutput(t *testing.T, output string, expectedStrings []string) {
	t.Helper()

	for _, str := range expectedStrings {
		if !strings.Contains(output, str) {
			t.Errorf("Expected output to contain %s", str)
		}
	}
}
func TestIntegration(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		out             *bytes.Buffer
		expectedStrings []string
		expectedErr     error
	}{
		{"Test `current`",
			[]string{"current", "-city", "berlin"},
			bytes.NewBuffer(nil),
			[]string{"Current weather for", "Feels like", "Expect"},
			nil,
		},
		{"Test `weekly`",
			[]string{"weekly", "-city", "berlin"},
			bytes.NewBuffer(nil),
			[]string{"Not implemented yet"},
			nil,
		},

		{"Test `help`",
			[]string{"help"},
			bytes.NewBuffer(nil),
			[]string{"$ current -city=Berlin", "weekly  -city=Toronto"},
			nil,
		},

		{"Test empty city flag",
			[]string{"weekly", "-city", ""},
			bytes.NewBuffer(nil),
			[]string{""},
			ErrorCityMissing,
		},
		{"Test unsupported command",
			[]string{"hourly", "-city", "berlin"},
			bytes.NewBuffer(nil),
			[]string{""},
			ErrorUnknownCommand,
		},
		{"Test insufficient args",
			[]string{},
			bytes.NewBuffer(nil),
			[]string{""},
			ErrorInsufficientArgs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			err := setupCLI(tt.args, weatherAPI, tt.out)
			assertError(t, tt.expectedErr, err)
			stringOutput := tt.out.String()
			assertConsoleOutput(t, stringOutput, tt.expectedStrings)
		})
	}
}

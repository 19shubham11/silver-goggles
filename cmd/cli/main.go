package main

import (
	config "19shubham11/weather-cli/config"
	weather "19shubham11/weather-cli/pkg/weather"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	appConfig := config.LoadAppConfig()
	weatherAPI := weather.WeatherAPI{
		Conf: appConfig,
	}
	var out io.Writer = os.Stdout

	if err := setupCLI(os.Args[1:], weatherAPI, out); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setupCLI(args []string, weatherAPI weather.API, out io.Writer) error {
	if len(args) < 1 {
		fmt.Println("--help")
		fmt.Println("$ current -city=<cityName>")
		fmt.Println("$ weekly -city=<cityName>")
		return errors.New("pass a valid option - `current` or `weekly`")
	}

	cmds := []Runner{
		NewWeatherCommand(CommandCurrentWeather, weatherAPI, out),
		NewWeatherCommand(CommandWeeklyWeather, weatherAPI, out),
	}

	subCommand := args[0]

	for _, cmd := range cmds {
		if cmd.Name() == subCommand {
			err := cmd.Init(args[1:])
			if err != nil {
				return err
			}
			return cmd.Run()
		}
	}
	return fmt.Errorf("Unknown subcommand: %s", subCommand)
}

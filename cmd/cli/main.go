package main

import (
	config "19shubham11/weather-cli/config"
	weather "19shubham11/weather-cli/pkg/weather"
	"errors"
	"fmt"
	"os"
)

func main() {
	appConfig := config.LoadAppConfig()
	weatherAPI := weather.WeatherAPI{
		Conf: appConfig,
	}

	if err := root(os.Args[1:], weatherAPI); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func root(args []string, weatherAPI weather.API) error {

	if len(args) < 1 {
		fmt.Println("--help")
		fmt.Println("$ current -city=<cityName>")
		fmt.Println("$ weekly -city=<cityName>")
		return errors.New("You must pass a valid argument - `current` or `weekly`")
	}

	cmds := []Runner{
		NewWeatherCommand(OptionCurrent, weatherAPI),
		NewWeatherCommand(OptionWeekly, weatherAPI),
	}

	subCommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subCommand {
			err := cmd.Init(os.Args[2:])
			if err != nil {
				return err
			}
			return cmd.Run()
		}
	}
	return fmt.Errorf("Unknown subcommand: %s", subCommand)
}

package main

import (
	config "19shubham11/weather-cli/config"
	weather "19shubham11/weather-cli/pkg/weather"
	"fmt"
	"io"
	"os"
)

func main() {
	appConfig := config.LoadAppConfig()
	weatherAPI := weather.OpenWeatherAPI{
		Conf: appConfig,
	}

	var out io.Writer = os.Stdout

	if err := setupCLI(os.Args[1:], weatherAPI, out); err != nil {
		// fmt.Println(err)
		os.Exit(0)
	}
}

func setupCLI(args []string, weatherAPI weather.API, out io.Writer) error {
	if len(args) < 1 {
		fmt.Fprintln(out, "insufficient arguments, see help ")
		return ErrorInsufficientArgs
	}

	cmds := []Runner{
		NewWeatherCommand(CommandHelp, nil, out),
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

	fmt.Fprintf(out, "unknown subcommand: %s, check help for usage", subCommand)

	return ErrorUnknownCommand
}

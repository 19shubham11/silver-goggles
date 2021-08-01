package main

import (
	weather "19shubham11/weather-cli/pkg/weather"
	"errors"
	"flag"
	"fmt"
)

const (
	OptionCurrent = "current"
	OptionWeekly  = "weekly"
)

var ErrorCityMissing = errors.New("city missing")

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

type WeatherCommand struct {
	fs   *flag.FlagSet
	city string
	api  weather.API
}

func NewWeatherCommand(commandName string, weatherAPI weather.API) *WeatherCommand {
	w := &WeatherCommand{
		fs:  flag.NewFlagSet(commandName, flag.ExitOnError),
		api: weatherAPI,
	}
	w.fs.StringVar(&w.city, "city", "", "name of the city")
	return w
}

func (w *WeatherCommand) Init(args []string) error {
	err := w.fs.Parse(args)
	if w.city == "" {
		fmt.Println("--help")
		fmt.Println(fmt.Sprintf("$ %s -city=london", w.fs.Name()))
		return ErrorCityMissing
	}
	return err
}

func (w *WeatherCommand) Name() string {
	return w.fs.Name()
}

func (w *WeatherCommand) Run() error {
	switch w.fs.Name() {
	case OptionCurrent:
		weather, err := w.api.GetCurrentWeather(w.city)
		if err != nil {
			fmt.Println("error!", err)
			return err
		}
		fmt.Println(weather)
	case OptionWeekly:
		fmt.Println("Not implemented yet!")
		// return nil
	}
	return nil
}

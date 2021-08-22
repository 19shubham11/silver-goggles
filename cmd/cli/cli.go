package main

import (
	weather "19shubham11/weather-cli/pkg/weather"
	"errors"
	"flag"
	"fmt"
	"io"
)

const (
	CommandCurrentWeather = "current"
	CommandWeeklyWeather  = "weekly"
	CommandHelp           = "help"
)

var (
	ErrorCityMissing      = errors.New("city missing")
	ErrorUnknownCommand   = errors.New("unknown command")
	ErrorInsufficientArgs = errors.New("insufficient args")
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

type WeatherCommand struct {
	fs     *flag.FlagSet
	city   string
	api    weather.API
	output io.Writer
}

func NewWeatherCommand(commandName string, weatherAPI weather.API, out io.Writer) *WeatherCommand {
	w := &WeatherCommand{
		fs:     flag.NewFlagSet(commandName, flag.ExitOnError),
		api:    weatherAPI,
		output: out,
	}
	if commandName != CommandHelp {
		w.fs.StringVar(&w.city, "city", "", "name of the city")
	}
	return w
}

func (w *WeatherCommand) Init(args []string) error {
	err := w.fs.Parse(args)
	if w.Name() != CommandHelp && w.city == "" {
		fmt.Fprintln(w.output, "-help")
		fmt.Fprintln(w.output, fmt.Sprintf("$ %s -city=london", w.fs.Name()))
		return ErrorCityMissing
	}
	return err
}

func (w *WeatherCommand) Name() string {
	return w.fs.Name()
}

func (w *WeatherCommand) Run() error {
	switch w.Name() {
	case CommandHelp:
		return w.getHelp()
	case CommandCurrentWeather:
		return w.getCurrentWeather()
	case CommandWeeklyWeather:
		return w.getWeeklyWeather()
	}
	return nil
}

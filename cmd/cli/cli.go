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
)

var ErrorCityMissing = errors.New("city missing")

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
	w.fs.StringVar(&w.city, "city", "", "name of the city")
	return w
}

func (w *WeatherCommand) Init(args []string) error {
	err := w.fs.Parse(args)
	if w.city == "" {
		fmt.Fprintln(w.output, "--help")
		fmt.Fprintln(w.output, fmt.Sprintf("$ %s -city=london", w.fs.Name()))
		return ErrorCityMissing
	}
	return err
}

func (w *WeatherCommand) Name() string {
	return w.fs.Name()
}

func (w *WeatherCommand) Run() error {
	switch w.fs.Name() {
	case CommandCurrentWeather:
		err := w.getCurrentWeather()
		return err
	case CommandWeeklyWeather:
		fmt.Fprintln(w.output, "Not implemented yet!")
		return nil
	}
	return nil
}

func (w *WeatherCommand) getCurrentWeather() error {
	weather, err := w.api.GetCurrentWeather(w.city)
	if err != nil {
		fmt.Println("error!", err)
		return err
	}
	// fmt.Printf("%+v\n", weather)

	fmt.Fprintf(w.output, "Current weather for %s\n", w.city)
	fmt.Fprintf(w.output, "Feels like %.2f°C\n", weather.Values.FeelsLike)
	fmt.Fprintf(w.output, "Expect %s\n", weather.Weather[0].Description)
	return nil
}

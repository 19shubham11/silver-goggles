package main

import "fmt"

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

func (w *WeatherCommand) getWeeklyWeather() error {
	fmt.Fprintln(w.output, "Not implemented yet!")
	return nil
}

func (w *WeatherCommand) getHelp() error {
	fmt.Fprintln(w.output, "usage")
	fmt.Fprintln(w.output, "$ current -city=Berlin")
	fmt.Fprintln(w.output, "$ weekly  -city=Toronto")
	return nil
}

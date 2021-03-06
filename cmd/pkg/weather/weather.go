package weather

import (
	"encoding/json"
	"errors"
	"net/http"

	"19shubham11/weather-cli/cmd/config"
	"19shubham11/weather-cli/cmd/internal/httpclient"
)

type OpenWeatherAPI struct {
	Conf *config.Config
}

func (w OpenWeatherAPI) GetCurrentWeather(cityName string) (*CurrentWeather, error) {
	url := w.Conf.WeatherURL

	queryParmas := map[string]string{
		"q":     cityName,
		"units": "metric",
		"appid": w.Conf.APIKey,
	}

	res, err := httpclient.Get(url, nil, queryParmas)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		currentWeather := &CurrentWeather{}
		err = json.NewDecoder(res.Body).Decode(currentWeather)

		if err != nil {
			return nil, err
		}

		return currentWeather, nil
	}

	return nil, errors.New("openweather error")
}

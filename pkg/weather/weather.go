package weather

import (
	config "19shubham11/weather-cli/config"
	httpClient "19shubham11/weather-cli/internal/httpclient"
	"encoding/json"
	"errors"
)

type WeatherAPI struct {
	Conf *config.Config
}

func (w WeatherAPI) GetCurrentWeather() (*CurrentWeather, error) {

	url := w.Conf.WeatherURL

	res, err := httpClient.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()

		currentWeather := &CurrentWeather{}
		err = json.NewDecoder(res.Body).Decode(currentWeather)
		if err != nil {
			return nil, err
		}
		return currentWeather, nil
	} else {
		return nil, errors.New("openweather error")
	}
}

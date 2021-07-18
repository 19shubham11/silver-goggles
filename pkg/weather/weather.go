package weather

import (
	httpClient "19shubham11/weather-cli/internal/httpclient"
	"encoding/json"
)

func GetCurrentWeather() (*CurrentWeather, error) {
	// move this to conf
	url := "http://api.openweathermap.org/data/2.5/weather?q=berlin&appid=**&units=metric"

	res, err := httpClient.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// better error handling and more tests
	if res.StatusCode == 200 {
		currentWeather := &CurrentWeather{}
		err = json.NewDecoder(res.Body).Decode(currentWeather)
		if err != nil {
			return nil, err
		}
		return currentWeather, nil
	}

	return nil, err
}

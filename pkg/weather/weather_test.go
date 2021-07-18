package weather

import (
	config "19shubham11/weather-cli/config"
	httpClient "19shubham11/weather-cli/internal/httpclient"
	"19shubham11/weather-cli/test/mocks"
	"testing"
)

var weatherAPI WeatherAPI

func init() {
	httpClient.Client = &mocks.MockClient{}

	mockConf := &config.Config{
		WeatherURL: "https://some-url.org",
		ApiKey:     "hahaLoL",
	}
	weatherAPI = WeatherAPI{
		Conf: mockConf,
	}
}

func assertSuccessfulResponse(t *testing.T, resp *CurrentWeather, err error, temperature float64) {
	t.Helper()

	if resp == nil {
		t.Fatalf("Empty response!")
	}

	if err != nil {
		t.Errorf("Expected nil, got err %v", err)
	}

	if resp.Values.Temp != temperature {
		t.Errorf("Expected %f, got %f", resp.Values.Temp, 254.35)
	}
}

func assertErrorResponse(t *testing.T, resp *CurrentWeather, err error) {
	t.Helper()

	if resp != nil {
		t.Errorf("Expected nil, got resp %v", resp)
	}

	if err.Error() != "openweather error" {
		t.Errorf("unexpected error")
	}
}

func TestGetCurrentWeather(t *testing.T) {
	t.Run("Should return proper response when remote server returns 200", func(t *testing.T) {
		json := `{"weather":[{"id":801,"main":"Clear","description":"clear sky","icon":"01d"}],"main":{"temp":254.35,"feels_like":25.46,"temp_min":24.57,"temp_max":26.15,"pressure":1015,"humidity":58}}`

		mocks.GetDoFunc = mocks.MockHTTPRequest(json, 200)

		resp, err := weatherAPI.GetCurrentWeather()
		assertSuccessfulResponse(t, resp, err, 254.35)
	})

	t.Run("return proper error", func(t *testing.T) {
		json := ""
		mocks.GetDoFunc = mocks.MockHTTPRequest(json, 400)

		resp, err := weatherAPI.GetCurrentWeather()
		assertErrorResponse(t, resp, err)
	})
}

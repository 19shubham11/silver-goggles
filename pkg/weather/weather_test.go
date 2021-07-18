package weather

import (
	httpClient "19shubham11/weather-cli/internal/httpclient"
	"19shubham11/weather-cli/test/mocks"
	"testing"
)

func init() {
	httpClient.Client = &mocks.MockClient{}
}

func assertSuccessfulResponse(t *testing.T, resp *CurrentWeather, err error, temperature float64) {
	t.Helper()

	if resp == nil {
		t.Fatalf("Empty response!")
	}

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if resp.Values.Temp != temperature {
		t.Errorf("Expected %f, got %f", resp.Values.Temp, 254.35)
	}
}

func TestGetCurrentWeather(t *testing.T) {
	json := `{"weather":[{"id":801,"main":"Clear","description":"clear sky","icon":"01d"}],"main":{"temp":254.35,"feels_like":25.46,"temp_min":24.57,"temp_max":26.15,"pressure":1015,"humidity":58}}`

	mocks.GetDoFunc = mocks.MockHTTPRequest(json, 200)

	resp, err := GetCurrentWeather()
	assertSuccessfulResponse(t, resp, err, 254.35)
}

package weather

type API interface {
	GetCurrentWeather(cityName string) (*CurrentWeather, error)
}

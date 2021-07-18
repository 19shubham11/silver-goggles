package weather

type weather []struct {
	Conditions  string `json:"main"`
	Description string `json:"description"`
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type CurrentWeather struct {
	Weather weather `json:"weather"`
	Values  main    `json:"main"`
}

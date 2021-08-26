package weather

type Description []struct {
	Summary     string `json:"main"`
	Description string `json:"description"`
}

type Temperature struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type CurrentWeather struct {
	Weather Description `json:"weather"`
	Values  Temperature `json:"main"`
}

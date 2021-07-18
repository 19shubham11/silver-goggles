package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	WeatherURL string
	ApiKey     string
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = fallback
	}
	return value
}

func LoadAppConfig() *config {

	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/config/config.json")
	if err != nil {
		fmt.Println("err", err)
		panic("error loading app config")
	}

	appConfig := &config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(appConfig)
	if err != nil {
		panic("error decoding app config")
	}

	// assign env variables to config
	appConfig.ApiKey = getEnv("OPENWEATHER_KEY", "testKey")

	return appConfig
}

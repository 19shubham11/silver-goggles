package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
)

type Config struct {
	WeatherURL string
	APIKey     string
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = fallback
	}

	return value
}

func LoadAppConfig() *Config {
	// Relative on runtime DIR:
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("error resolving directory")
	}

	dir := path.Join(path.Dir(b))

	file, err := os.Open(dir + "/config.json")
	if err != nil {
		fmt.Println("err", err)
		panic("error loading app config")
	}

	appConfig := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(appConfig)

	if err != nil {
		panic("error decoding app config")
	}

	// assign env variables to config
	appConfig.APIKey = getEnv("OPENWEATHER_KEY", "testKey")

	return appConfig
}

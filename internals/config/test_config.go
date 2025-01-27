package config

import (
	"os"
)

type TestConfig struct {
	DB_DRIVER string `map_structure:"DB_DRIVER"`
	DB_SOURCE string `map_structure:"DB_SOURCE"`
}

func LoadTestConfig() (config Config, err error) {

	config.DB_DRIVER = os.Getenv("DB_DRIVER")
	config.DB_SOURCE = os.Getenv("DB_SOURCE")

	return
}

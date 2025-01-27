package config

import "github.com/spf13/viper"

type TestConfig struct {
	DB_DRIVER string `map_structure:"DB_DRIVER"`
	DB_SOURCE string `map_structure:"DB_SOURCE"`
}

func LoadTestConfig() (config Config, err error) {

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

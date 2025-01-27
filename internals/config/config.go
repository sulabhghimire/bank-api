package config

import "github.com/spf13/viper"

// Config contains all the configuration for the application
// The values are read by viper from configuration files or environmental variables
type Config struct {
	DB_DRIVER      string `map_structure:"DB_DRIVER"`
	DB_SOURCE      string `map_structure:"DB_SOURCE"`
	SERVER_ADDRESS string `map_structure:"SERVER_ADDRESS"`
}

// Loads the config from the config file
func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

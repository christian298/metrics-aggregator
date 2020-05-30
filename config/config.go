package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Db     DbConfig
}

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	Name     string
	URL      string
	User     string
	Password string
}

func ReadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file")
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		fmt.Printf("Error decoding config file")
	}

	return config
}

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Environment string
	Server      ServerConfig
	Db          DbConfig
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
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file")
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		fmt.Printf("Error decoding config file")
	}

	fmt.Println(config)

	return config
}

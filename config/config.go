package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error

	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config/")
	// Path for testing
	config.AddConfigPath("../config/")

	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("Error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}

package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string `mapstructure:"connection_string"`
	Port             string `mapstructure:"port"`
}

var Configuration *Config

func LoadConfig() {
	log.Println("Loading Configurations...")

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	err = viper.Unmarshal(&Configuration)
	if err != nil {
		log.Fatalln(err)
	}
}

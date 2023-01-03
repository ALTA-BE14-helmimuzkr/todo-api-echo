package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
	DBNAME string
}

func InitConfig() *AppConfig {
	appConfig := AppConfig{}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Load config
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Read config failed", err)
		return nil
	}

	// Parsing config dari hasil read(json) ke struct app config
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Println("Unmarshal config failed", err)
		return nil
	}

	return &appConfig
}

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

	viper.AddConfigPath("./.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Gagal membaca config", err)
		return nil
	}

	return &appConfig
}

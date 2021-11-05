package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppHost    string `mapstructure:"app_host"`
	AppPort    int16  `mapstructure:"app_port"`
	PgHost     string `mapstructure:"pg_host"`
	PgUsername string `mapstructure:"pg_username"`
	PgPassword string `mapstructure:"pg_password"`
	PgDbname   string `mapstructure:"pg_dbname"`
	PgPort     int16  `mapstructure:"pg_port"`
}

func GetAppConfig() *AppConfig {
	var err error

	var defaConfig AppConfig
	defaConfig.AppHost = "localhost"
	defaConfig.AppPort = 5000
	defaConfig.PgHost = "localhost"
	defaConfig.PgUsername = "postgres"
	defaConfig.PgPassword = ""
	defaConfig.PgDbname = "postgres"
	defaConfig.PgPort = 5432

	// Fetch configuration file
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error read file config, app run default config")
		return &defaConfig
	}

	var finalConfig AppConfig
	err = viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Fatal("Error unmarshal file config, app run default config")
		return &defaConfig
	}

	return &finalConfig
}

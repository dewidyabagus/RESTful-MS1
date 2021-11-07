package config

import (
	"log"
	"sync"

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

type AppSecretKey struct {
	JwtSecretKey string `mapstructure:"jwt_secret_key"`
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

func GetJWTSecretKey() string {
	var mtx = &sync.Mutex{}

	mtx.Lock()
	defer mtx.Unlock()

	var appSecret AppSecretKey

	appSecret.JwtSecretKey = "default_example_config@example.com"

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return appSecret.JwtSecretKey
	}

	err = viper.Unmarshal(&appSecret)
	if err != nil {
		return appSecret.JwtSecretKey
	}

	return appSecret.JwtSecretKey
}

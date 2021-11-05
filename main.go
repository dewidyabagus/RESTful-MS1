package main

import (
	// Module configuration file
	"RESTfulMS1/config"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {
	strConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.PgHost, config.PgUsername, config.PgPassword, config.PgDbname, config.PgPort,
	)

	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	// Get configuration file
	config := config.GetAppConfig()

	_ = newDatabaseConnection(config)

	fmt.Println("Sukses")
}

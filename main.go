package main

import (
	// Module configuration file
	"RESTfulMS1/config"

	// Migration
	"RESTfulMS1/modules/migration"

	// API
	"RESTfulMS1/api"

	// User
	userController "RESTfulMS1/api/v1/user"
	userService "RESTfulMS1/business/user"
	userRepository "RESTfulMS1/modules/user"

	"fmt"

	echo "github.com/labstack/echo/v4"
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

	migration.TableMigration(db)

	return db
}

func main() {
	// Get configuration file
	config := config.GetAppConfig()

	// Create new session database
	dbConnection := newDatabaseConnection(config)

	// Initiate user repository
	userRepo := userRepository.NewRepository(dbConnection)

	// Initiate user service
	userSvc := userService.NewService(userRepo)

	// Initiate user controller
	userCtr := userController.NewController(userSvc)

	// Initiate echo web framework
	e := echo.New()

	// Initiate routes
	api.RegisterRouters(e, userCtr)

	// start echo
	e.Start(":8000")
}

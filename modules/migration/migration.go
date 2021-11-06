package migration

import (
	"gorm.io/gorm"

	"RESTfulMS1/modules/user"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}

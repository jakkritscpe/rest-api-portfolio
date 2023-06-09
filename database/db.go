package database

import (
	"log"
	"os"

	"github.com/jakkritscpe/rest-api-portfolio/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := os.Getenv("DNS_DATABASE")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		log.Println("Connect database Success.")
	}

	// Database Auto Migrate.
	Db.AutoMigrate(
		&models.User{},
		&models.Tools{},
		&models.CategoryTools{},
		&models.Skills{},
		&models.Projects{},
	)

}

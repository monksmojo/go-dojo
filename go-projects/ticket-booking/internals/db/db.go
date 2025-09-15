package db

import (
	"log"
	"os"

	"github.com/monksmojo/go-dojo/go-projects/open-talk/internals/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DATABASE_DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	DB = database
}

func BootStrapDB() {
	DB.AutoMigrate(&models.Venue{}, &models.Event{}, &models.Talk{}, &models.Speaker{})
}

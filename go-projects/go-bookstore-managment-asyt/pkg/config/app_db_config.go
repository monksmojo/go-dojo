package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/monks_mojo/go-dojo/go-projects/go-bookstore-managment-asyt/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func loadEnv() {
	godotenv.Load("../../.env")
}

func connectToDB() {
	dsn := getDbCredString(os.Getenv("PORT_IP"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = dbConnection
}

func GetDB() *gorm.DB {
	return db
}

type dbCred struct {
	portIp       string
	user         string
	password     string
	databaseName string
}

func getDbCredString(portIp, user, password, databaseName string) string {
	dbCredInstance := dbCred{
		portIp:       portIp,
		user:         user,
		password:     password,
		databaseName: databaseName,
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCredInstance.user, dbCredInstance.password, dbCredInstance.portIp, dbCredInstance.databaseName)
}

func bootStrapDB() {
	initBookDB()
}

func initBookDB() {
	db.AutoMigrate(&models.Book{})
}

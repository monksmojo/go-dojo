package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToMySqlDB() {
	var err error
	user:=os.Getenv("USER")
	password:=os.Getenv("PASSWORD")
	address:=os.Getenv("ADDRESS")
	dbName:=os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, address, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the the mysql database", err)
	}
}

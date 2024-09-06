package initializers

import (
	"log"
	"os"
	"strconv"

	"github.com/monksmojo/go-dojo/go-projects/go-gin-crud-api-by-robby/models"
)

func MigrateModelsToDB() {
	if migrate, _ := strconv.ParseBool(os.Getenv("AUTO_MIGRATE")); migrate {
		log.Println("migrating models to database")
		DB.AutoMigrate(&models.Employee{})
	} else {
		log.Println("models already migrated")

	}
}

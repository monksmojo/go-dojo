package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/monksmojo/go-dojo/go-projects/go-gin-crud-api-by-robby/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToMySqlDB()
	initializers.MigrateModelsToDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(os.Getenv("PORT")) // listen and serve on 0.0.0.0:9000
}

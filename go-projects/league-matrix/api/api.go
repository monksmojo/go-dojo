package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/monksmojo/go-dojo/go-project/league-matrix/models"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// The GetSwaggerUrl function returns a function that generates the Swagger URL
//  based on the environment variable PORT.
func GetSwaggerUrl() func(*ginSwagger.Config) {
	url := fmt.Sprintf("http://localhost:%v/swagger/doc.json", os.Getenv("PORT"))
	log.Printf("swagger url: %v \n", url)
	return ginSwagger.URL(url)
}

// The function `ErrorResponse` converts the error to an HTTP status code,and specific message 
// returns a JSON response with the status code and message./
func ErrorResponse(ginContext *gin.Context, err error) {
	code, msg := toHTTPStatusCode(err)
	ginContext.JSON(code, gin.H{"status": code, "message": msg})
}

// The SuccessResponse function sends a successful response with data in JSON format
func SuccessResponse(ginContext *gin.Context, data interface{}) {
	ginContext.JSON(200, gin.H{"status": 200, "message": "request processed successfully", "data": data})
}

func toHTTPStatusCode(err error) (int, string) {
	switch {
	case models.ErrParseError(err):
		return http.StatusInternalServerError, "error while fetching data from the csv file"
	case models.ErrNumError(err):
		return http.StatusInternalServerError, "csv file contains bad characters"
	case errors.Is(err, models.ErrMissingArgument):
		return http.StatusBadRequest, "bad request missing argument"
	case errors.Is(err, models.ErrInvalidMessageType):
		return http.StatusBadRequest, "bad request"
	case errors.Is(err, models.ErrNotFound):
		return http.StatusNotFound, "requested url not found"
	case errors.Is(err, models.ErrNoSuchFile):
		return http.StatusBadRequest, "file not present or incorrect file path"
	default:
		return http.StatusInternalServerError, "something went wrong"
	}
}

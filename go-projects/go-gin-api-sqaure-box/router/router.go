package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monksmojo/go-dojo/go-project/go-gin-api-square-box/api"
	"github.com/monksmojo/go-dojo/go-project/go-gin-api-square-box/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// The RegisterRoutes function sets up routes for handling matrix operations
func RegisterRoutes() http.Handler {
	r := gin.Default()
	handler := handler.FetchHandlers()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, api.GetSwaggerUrl()))
	r.POST("/echo", handler.MatrixHandler.Echo)
	r.POST("/invert", handler.MatrixHandler.Invert)
	r.POST("/flatten", handler.MatrixHandler.Flatten)
	r.POST("/sum", handler.MatrixHandler.Sum)
	r.POST("/multiply", handler.MatrixHandler.Multiply)
	return r
}

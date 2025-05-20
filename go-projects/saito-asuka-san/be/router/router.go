package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() http.Handler {
	r := initRouter()
	return r
}

func initRouter() http.Handler {
	r := gin.Default()
	return r
}

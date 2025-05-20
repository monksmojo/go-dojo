package routes

import (
	"net/http"

	"github.com/monks_mojo/go-dojo/go-projects/go-bookstore-managment-asyt/pkg/controllers"
)

func RegisterBookRoutes(serveMux *http.ServeMux) {
	serveMux.HandleFunc("POST /book/",controllers.CreateBook)
	serveMux.HandleFunc("GET /book/",controllers.GetAllBooks)
	serveMux.HandleFunc("GET /book/{bookId}",controllers.GetBookById)
	serveMux.HandleFunc("PUT /book/{bookId}",controllers.UpdateBookById)
	serveMux.HandleFunc("DELETE /book/{bookId}",controllers.DeleteBookById)
}

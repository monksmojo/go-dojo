package main

import (
	"log"
	"net/http"

	"github.com/monksmojo/go-dojo/go-projects/tutorialEdge-serveMux-http-request-in-go/handler"
)

func RequestLoggerMiddleWare(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func RequireAuthMIddleWare(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Qwerty@123" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// The function `MiddlewareChain` takes a slice of middleware functions and returns a new middleware
// function that chains them together.

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares []Middleware) Middleware {
	return func(h http.Handler) http.HandlerFunc {
		for _, middleware := range middlewares {
			h = middleware(h)
		}
		return h.ServeHTTP
	}
}

type ApiServer struct {
	addr     string
	serveMux *http.ServeMux
}

func (apiServer *ApiServer) Run() error {
	apiServer.serveMux.HandleFunc("Get /users/{userId}", handler.GetUserById)

	middlewareChain := MiddlewareChain([]Middleware{RequireAuthMIddleWare, RequestLoggerMiddleWare})

	server := http.Server{
		Addr:    apiServer.addr,
		Handler: middlewareChain(apiServer.serveMux),
	}
	log.Printf("starting the server at port %v", apiServer.addr)
	return server.ListenAndServe()
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr:     addr,
		serveMux: http.NewServeMux(),
	}
}

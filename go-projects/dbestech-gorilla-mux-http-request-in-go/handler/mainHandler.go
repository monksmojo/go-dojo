package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerProvider() *routerServer {
	routerServerInstance := routerServer{
		router: mux.NewRouter(),
	}
	routerServerInstance.http = &http.Server{Addr: "localhost:" + "8080", Handler: routerServerInstance.router}

	return &routerServerInstance
}

func RouterMapper(rs *routerServer) {
	rs.mapRouter()
}

type routerServer struct {
	router *mux.Router
	http   *http.Server
}

func (rs *routerServer) mapRouter() {
	rs.router.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is the response from the server")
	}).Methods("GET")
}

func (rs *routerServer) StartServer() error {
	err := rs.http.ListenAndServe()
	if err != nil {
		return err
	}
	fmt.Println("server started at port 8080")
	return nil
}

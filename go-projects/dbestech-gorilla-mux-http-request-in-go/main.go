package main

import (
	"fmt"
	"strings"

	"github.com/monksmojo/go-dojo/go-projects/dbestech-http-request-in-go/handler"
)

func main() {
	fmt.Println(strings.Repeat("#", 20), "welcome to", strings.Repeat("#", 20))
	fmt.Println(strings.Repeat("#", 20), "Http Request in go", strings.Repeat("#", 20))
	routerServer := handler.ServerProvider()
	handler.RouterMapper(routerServer)
	routerServer.StartServer()

}

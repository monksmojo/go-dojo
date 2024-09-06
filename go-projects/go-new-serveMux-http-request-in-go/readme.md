<div align=center>tutorialEdge-serveMux-http-request-in-go</div>

#### This project repository gives the idea of how we can use net/http package

#### of golang 1.22 which now has the capability to create a multiplexer that can act as a router as well as a net http server

#### We will understand the package by creating a _CRUD BACKEND API PROJECT_

#### following list of articles and youtube video lectures were consumed to create this repository

1. [Building REST APIs in Go 1.22 - New Features](https://youtu.be/tgLvIghsJFo) by _Elliot Forbes | TutorialEdge.net_
2. [How to create your Golang APIs from now on with v1.22](https://youtu.be/npzXQSL4oWo?list=PLYEESps429vqQ98y_zjyERFQR1cyvBNzA) by Tiago

> a heartiest thanks to all the creators of the video lectures and articles for the content ♥

### by following the tutorial by the elliot Forbes from the TutorialEdge.net

### we got the basic knowledge of the ServeMux() by go 1.22

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("GO BACKEND PROJECT USING NEW SERVE-MUX AND NET/HTTP OF GO1.22")
	muxRouter := http.NewServeMux()

	muxRouter.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "return all comments \n")
	})

	muxRouter.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "returning single comment with id %s \n", id)
	})

	muxRouter.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "post a new comment \n")
	})

	err := http.ListenAndServe("localhost:8080", muxRouter)
	if err != nil {
		log.Fatal("difficulty in starting the server")
	}
}


```

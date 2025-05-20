package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err; %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request successful")
	email := r.FormValue("email")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	fmt.Fprintf(w, "Name = %s\n", email)
	fmt.Fprintf(w, "First Name = %s\n", firstName)
	fmt.Fprintf(w, "Last Name = %s\n", lastName)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Started the server at the port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

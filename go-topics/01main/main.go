package main

import "fmt"

func main() {
	fmt.Println("Yo Hello Hello!")
	Greet("")
	Greet("Foo")
}

// every program has to have  a main function
//  it tells go where the program start it starts from the main function
// the main function should be in main package
// the function start with the `func` keyword

func Greet(name string) string {
	if name == "" {
		return "Greetings ! World \n Welcome To Go Topics, By Monks Mojo"
	}
	return fmt.Sprintf("Greetings ! %s \n Welcome To Go Topics, By Monks Mojo", name)
}

// go run . -- . means run in this directory
//  go run <file_name>.go
//  go run compiles and build and store the binary program in some temp directory/folder and after executing the binary it removes it from the temp directory


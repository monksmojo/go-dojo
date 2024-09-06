package main

import "fmt"

func main() {
	fmt.Println("Yo Hello Hello!")
	Greet("Foo")
}

func Greet(name string) string {
	if name == "" {
		return fmt.Sprint("Greetings ! World \n Welcome To Go Topics, By Monks Mojo")
	}
	return fmt.Sprintf("Greetings ! %s \n Welcome To Go Topics, By Monks Mojo", name)
}

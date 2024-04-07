package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("welcome to the closure in the go")
	// go closure is just a way to execute the nested function outside outer function,
	// even after the outer function is closed.

	// closure mainly use anonymous function
	// calling the greet function

	// the nested function now acts as a closure that closes the outer scope variable of outer function within its scope even after the outer function is executed.
	// so Go closure is a nested function that allows us to store and access variables of the outer function even after the outer function is closed.
	fmt.Println("calling the greet function")
	msg := greet("tom")
	fmt.Println("greet function exited")
	fmt.Println(msg())

	/**
	 a go closure helps in data isolation
	 a new closure is returned every time we call the outer function. Here, each returned closure is independent of one another, and the change in one won't affect the other.
	**/

	// printing 10 odd numbers using closure
	fmt.Println("printing 10 odd numbers using closure")
	oddNumber := getNumber(1)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d :: %d \n", i, oddNumber())
	}

	// printing 10 even numbers using closure
	fmt.Println("printing 10 even numbers using closure")
	evenNumber := getNumber(0)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d :: %d \n", i, evenNumber())
	}

	// evenNumber() closure is independent of oddNumber() closure

	/*
		A Closure Example from the free code camp
		A closure is a function that reference variables from outside its own function body
		the function may access and assigned to the reference variable.
	*/
	nameAggregator := concatenation()
	// nameAggregator() var nameAggregator func(string) string
	nameAggregator("Hi")
	nameAggregator("my")
	nameAggregator("name")
	nameAggregator("is")
	var concatenatedString string = nameAggregator("Paul")
	fmt.Println(concatenatedString)

}

func greet(name string) func() string {
	name = strings.ToUpper(name)
	fmt.Println("welcome greetings")
	// here we can see greet function returns an anonymous function
	return func() string {
		return fmt.Sprintf("Hey Hi ! %s", name)
	}
}

func getNumber(number int) func() int {
	return func() int {
		number = number + 2
		return number
	}
}

func concatenation() func(string) string {
	var doc string
	aggregator := func(word string) string {
		doc += word + " "
		return doc
	}
	return aggregator
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	var oneValue int = 23
	var onePtr *int = &oneValue
	var defaultPtr *string
	fmt.Println("default value of a pointer is", defaultPtr)
	// & is used to reference memory location of the variable
	// * is used to fetch the value store in the memory location pointer is holding
	// for example increase the value of oneValue by 5 using the pointer pointing to the one Value
	fmt.Println("what pointer actually holds", onePtr)
	fmt.Println("accessing value of the memory pointer points to", *onePtr)
	fmt.Println("value of oneValue before operation", oneValue)
	*onePtr = *onePtr + 5
	fmt.Println("value of oneValue after operation", oneValue)

}

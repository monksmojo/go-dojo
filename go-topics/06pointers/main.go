package main

import (
	"fmt"
	"strings"
)

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
	fmt.Println("accessing value of the memory pointer points to", onePtr)
	fmt.Println("value of oneValue before operation", *onePtr)
	*onePtr = *onePtr + 5
	fmt.Println("value of oneValue after operation", oneValue)

	// shallow copy
	x := 5
	y := x // this is we making the shallow copy of variable x
	y = 10 // which gets replaced here
	println(x, y)

	// as example of passing pointers to the function
	// with classic swap function
	a := 420
	b := 69

	fmt.Println("value of a and b before swap")
	fmt.Println("a value ", a)
	fmt.Println("b value ", b)

	swapTwoDigit(&a, &b)

	fmt.Println("inside function value of a and b")
	fmt.Println("a value ", a)
	fmt.Println("b value ", b)

	msg := "shoot what the heck bro!!"
	removeParity(&msg)
	fmt.Println("after removing the parity", msg)
}

func swapTwoDigit(aPtr, bPtr *int) {
	aPtr, bPtr = bPtr, aPtr
	fmt.Println("inside function value of a and b")
	fmt.Println("a value ", *aPtr)
	fmt.Println("b value ", *bPtr)
}

// the function takes a msg string reference remove all the cuss words from it
func removeParity(msg *string) {
	tempMsg := *msg
	tempMsg = strings.ReplaceAll(tempMsg, "dang", "****")
	tempMsg = strings.ReplaceAll(tempMsg, "shoot", "******")
	tempMsg = strings.ReplaceAll(tempMsg, "heck", "****")
	*msg = tempMsg
}

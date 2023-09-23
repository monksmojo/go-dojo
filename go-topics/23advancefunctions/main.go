package main

import "fmt"

func main() {
	/*
		FIRST CLASS FUNCTION
		a first class function is the that can be treated like any other value passed to the other function
		a first class function can act as a argument to the other function
	*/
	/*
		HIGHER ORDER FUNCTION
		a higher order function is the function which takes a first class function as an argument
		or returns a function as an argument
	*/
	fmt.Println("welcome to advance function")

	a := 10
	b := 20
	c := 30
	var r1 int = aggregate(a, b, c, add)
	var r2 int = aggregate(a, b, c, multiply)
	fmt.Printf("sum of %d + %d + %d is %d \n", a, b, c, r1)
	fmt.Printf("product of %d * %d * %d is %d \n", a, b, c, r2)

	/*
		FUNCTION CURRYING
		function currying is the practice of writing a function that takes a function (or functions) as inputs
		and return a new function
	*/
	doubleFunc := selfMath(add)
	squareFunc := selfMath(multiply)
	fmt.Printf("double of %d is %d \n",a,doubleFunc(a))
	fmt.Printf("square of %d is %d \n",c,squareFunc(c))

}

// The function "add" takes two integers as input and returns their sum.
// the add is a first order function to aggregate
func add(a, b int) int {
	return a + b
}

// The function "multiply" takes two integers as input and returns their product.
// the multiply is a first order function to aggregate
func multiply(a, b int) int {
	return a * b
}

// the aggregate() function is a higher order function
// The function "aggregate" takes three integers and an arithmetic function as parameters, and returns
// the result of applying the arithmetic function to the three integers in a specific order.
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// selfMath() is the currying function
// The selfMath function takes an arithmetic function as input and returns a new function that applies
// the arithmetic function to a single input value twice.
func selfMath(arithmetic func(int, int) int) func(int) int {
	return func(x int) (r int) {
		return arithmetic(x, x)
	}
}

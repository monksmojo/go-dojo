package main

import (
	"errors"
	"fmt"
	"math"
)

var theValue float64 = 57.89

func main() {

	fmt.Println("welcome to functions in go")
	greet()

	fmt.Println("calling add function")
	var a uint = 89
	var b uint = 78
	var c uint = 90
	add(a, b)

	fmt.Println("calling product function")
	result := product(a, b)
	fmt.Println(result)

	var x float64 = 90.0
	var y float64 = 90.0
	fmt.Println("calling division function")
	r, q := division(x, y)
	fmt.Printf("%v / %v = %v \n", x, y, q)
	fmt.Printf("remainder of %v / %v = %v \n", x, y, r)

	fmt.Println("calling adding of the three numbers function")
	l := addThree(a, b, c)
	fmt.Printf("addition of %d + %d + %d = %d \n", a, b, c, l)

	// creating functions in go that takes indefinite number of parameter args
	var slice = []uint{2, 4, 6, 8}
	u := sumOfSlice(slice...)
	fmt.Println("sum of slices is", u)

	// creating a function in go that takes another function as a parameter
	average := calculateAverage(sumOfSlice, a, b, c)
	fmt.Printf("average of (%d,%d ,%d) = %f \n", a, b, c, average)

	// creating anonymous function in go
	anonymous := func(i uint) string {
		return fmt.Sprintf("calling closure %d times", i)
	}

	for i := 1; i <= 5; i++ {
		// converting int to uint
		fmt.Println(anonymous(uint(i)))
	}

	// its better to pass an anonymous function in go
	// as the argument to the another function
	sumOfTwo := func(n1 int, n2 int) (r int) {
		r = n1 + n2
		return r
	}

	// calling a regular function and passing anonymous function as an argument
	squareIs := calculateSquare(sumOfTwo(10, 10))
	fmt.Println("square is", squareIs)

	// calling a regular function that returns the anonymous function
	cubeOfThree := sumOfThree(6, 3, 1)
	// here cubeOfThree is anonymous function returned by sumOfThree()
	fmt.Println("cube of 10 is", cubeOfThree())

	// in go arguments in the function are passed as values
	var v1 uint = 67
	fmt.Printf("value of v1 before the increment function call is = %v \n", v1)
	increment := func(value uint) {
		fmt.Println("inside increment value ")
		fmt.Println("value incremented", value+uint(1))
	}
	increment(v1)
	fmt.Printf("value of v1 after the increment function call is = %v \n", v1)

	// calling function in go with named returns and and using naked return to return the default values.
	var v2 int = 6
	var v3 int = 8
	fmt.Println("values to get double are", v2, v3)
	v2, v3 = incorrectDoubleTheValue(v2, v3)
	fmt.Println("incorrect values", v2, v3)

	// calling function in go with named returns.
	v2, v3 = 6, 8
	fmt.Println("values to get double are", v2, v3)
	v2, v3 = correctDoubleTheValue(v2, v3)
	fmt.Println("correct values", v2, v3)

	result2, err := divide(10.0, 0.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result2)
	}

	// calling the changeTheGlobalValue()
	fmt.Println("global value before the function call", theValue)
	changeTheGlobalValue()
	fmt.Println("global value after the function call", theValue)

	//recursion in go
	fact := factorial(5)
	fmt.Println("factorial of 5 using recursion in go is", fact)

}

/**
functions in go without function parameters and return
**/

func greet() {
	fmt.Println("Hello ! Hi There")
}

/*
*
functions in go with parameters and no return
*
*/
func add(n1 uint, n2 uint) {
	c := n1 + n2
	fmt.Printf("sum of %v + %v = %v \n", n1, n2, c)
}

/*
*
functions in go with parameters and single return
*
*/
func product(n1 uint, n2 uint) string {
	c := n1 * n2
	return fmt.Sprintf("sum of %v * %v = %v \n", n1, n2, c)
}

/*
*
functions in go with parameters and multiple return
*
*/
func division(n1 float64, n2 float64) (float64, float64) {
	d := n2 / n1
	//  In Go, the modulus operator (%) cannot be used with floating-point numbers
	// so will use math.MOD()
	r := math.Mod(n2, n1)
	return d, r
}

/**
When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
**/

func addThree(x, y, z uint) uint {
	return x + y + z
}

// a variadic function is a function of indefinite arbitrary
// in go a function can take n numbers of same type of values as an argument to the function

func sumOfSlice(args ...uint) uint {
	var sum uint = 0
	// the args are treated as slices of the defined type
	fmt.Printf("the values fed to the function are %v and its type is %T \n", args, args)

	for _, v := range args {
		sum += v
	}
	return sum
}

// in go we can pass function as the parameter
func calculateAverage(sumOfSlice func(slice ...uint) uint, slice ...uint) float64 {
	sum := sumOfSlice(slice...)
	return float64(sum) / float64(len(slice))
}

// function double the values which has named return
// and a naked return
func incorrectDoubleTheValue(a, b int) (x int, y int) {
	a *= 2
	b *= 2
	// since we are not explicitly returning anything
	// the naked return will return
	// default value of the named parameter (x,y)
	fmt.Println("not doubling the value and returning naked default value of named return")
	// naked returns are bad for readability
	return
}

func correctDoubleTheValue(a, b int) (x int, y int) {
	x = a * 2
	y = b * 2
	fmt.Println("correctly doubling the values")
	return x, y
}

// function with early return
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by zero")
	} else {
		return dividend / divisor, nil
	}
}

func changeTheGlobalValue() {
	fmt.Println("changing the global value")
	theValue = 10.0
	fmt.Println("the global value inside the function", theValue)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

// creating a regular function that takes anonymous function as an argument
func calculateSquare(num int) int {
	fmt.Println("taking anonymous function as an argument and returning the square")
	return num * num
}

func sumOfThree(n1, n2, n3 int) func() int {
	n4 := n1 + n2 + n3
	fmt.Println("anonymous function returned by sumOfThree()")
	return func() int {
		return n4 * n4 * n4
	}
}

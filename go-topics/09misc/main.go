package main

import "fmt"

func main() {
	// a single line comment
	/*
		a multi
		line
		comment
	*/
	// 3 ways of printing the messages in go
	name := "chunk solo"
	age := 25
	hobby := "Driving you crazing"
	salary := 75000.23

	
	// fmt.Print()
	fmt.Print("using Print() \n")
	fmt.Print("name is", name)
	fmt.Print("age is ", age)
	fmt.Print("hobby is ", hobby, "\n")
	// in fmt.Print() If we print multiple values and variables at once,
	// NO space is added between the values by default.

	// fmt.Println()
	// all the output messages are output to the separate lines
	// by default a space is given when multiple values are there
	fmt.Println("using Println()")
	fmt.Println("name is", name)
	fmt.Println("age is ", age)
	fmt.Println("hobby is ", hobby)

	// fmt.Printf()
	// it is use for string formatting
	// integer	%d
	// float	%g,%f
	// string	%s
	// bool		%t
	// any value %v
	// type of the variable %T
	// pointer %p
	fmt.Printf("Hi ! my name is %s and age is %d and salary is %g \n", name, age, salary)
	

}

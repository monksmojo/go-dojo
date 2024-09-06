package main

import "fmt"

func main() {
	fmt.Println("welcome to the conditional programming in the go")

	// only if statement
	var yourAge uint = 28
	var friendsAge uint = 12
	var peopleAge uint = 27
	if yourAge > 18 {
		fmt.Println("you are an adult")
	}

	// if condition with initialization
	if myAge := 20; myAge > 20 {
		fmt.Println("you are old enough")
	}

	// if and else condition
	if friendsAge > 18 {
		fmt.Println("my friend is an adult")
	} else {
		fmt.Println("my friend is not an adult")
	}

	// if else if else ladder
	if peopleAge > 18 && peopleAge <= 20 {
		fmt.Println(" you are an adult but can't rent a car")
	} else if peopleAge > 20 && peopleAge <= 25 {
		fmt.Println("you can rent a car")
	} else {
		fmt.Println("go home kid")
	}

	// nested if else
	var num int = 0
	if num > 0 {
		fmt.Println(num, "number is positive")
		if num%2 == 0 {
			fmt.Println(num, "number is even")
		} else {
			fmt.Println(num, "number is odd")
		}
	} else if num < 0 {
		fmt.Println(num, "number is negative")
	} else {
		fmt.Println(num, "number is zero itself")
	}

}

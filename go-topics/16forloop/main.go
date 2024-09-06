package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to for loops in go")
	// in go lang we don't have while loop or do while loop
	fmt.Println("write a program to print factorial of the number")

	// factorial of the number using for while loop
	fmt.Println("factorial of the number using standard for loop")
	fmt.Printf("factorial of %d is ", 5)
	var fac int = 1
	for i := 5; i >= 1; i -= 1 {
		fac *= i
	}
	fmt.Printf("= %d \n", fac)

	// factorial of the number using for while loop
	fmt.Println("factorial of the number using for while loop")
	var fact uint = 1
	var i uint = 5
	fmt.Printf("factorial of %d is ", i)
	for i >= 1 {
		fact *= i
		i -= 1
	}
	fmt.Printf("= %d \n", fact)

	// for loop on the range of arrays
	var num = [5]int{10, 20, 30, 40, 50}

	fmt.Println("for loop on the array with default index")
	for index := range num {
		fmt.Printf("at index %v array has value %v \n", index, num[index])
	}

	fmt.Println("for loop on the array with index and value")
	for index, value := range num {
		fmt.Printf("at index %v array has value %v \n", index, value)
	}

	fmt.Println("for loop on the array if we only need value")
	for _, value := range num {
		fmt.Printf("%v \n", value)
	}

	fmt.Println("using for loop as a while loop to print the multiplication table")
	var multiplier uint = 1
	for multiplier <= 10 {
		fmt.Printf("5 * %d = %d \n", multiplier, multiplier*5)
		multiplier += 1
	}

	fmt.Println("using for loop as do while to print count till 10")
	digit := 1
	for {
		if digit > 10 {
			break
		} else {
			fmt.Println(digit)
			digit += 1
		}
	}

	// go for range with string
	fmt.Println("for range with string")
	for index, value := range "grand prix" {
		fmt.Printf("s[%d] = %c . \t", index, value)
	}
	fmt.Println()

	var st1 string = "copper mesh"
	st1Slice := []rune(st1)
	fmt.Printf("converting string %v into %v \n", st1, st1Slice)
	for _, v := range st1Slice {
		fmt.Printf("%v : %v\n", v, string(v))
	}
	fmt.Println()

	// go for range with maps
	// maps are not ordered while traversing
	var studentMarks = map[string]float64{"M": 76.78, "I": 87.89, "K": 56.30, "E": 95.89}
	for k, v := range studentMarks {
		fmt.Printf(" [%v] :%v \n", k, v)
	}

	// for loop with break and continue
	// for loop with break
	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println(i)
	}

}

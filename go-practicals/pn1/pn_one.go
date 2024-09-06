package pn1

import (
	"fmt"
	"math"
)

// Practical1 Description
/**
@initialAuthor: monks_mojo
@lastUpdatedBY: monks_mojo
Practical One:
WAF which should print all the prime numbers up to the given nTh integer.
Including the nTh integer.

Test case sample:1
input: nth integer : 17
output: list of prime number till 17 [2,3,5,7,11,13,17]
**/

/*
The function "CalculatePrime" takes user input for the value of "nThInteger" and then calls the
"printPrimeTill" function to print all prime numbers up to and including the "nThInteger".
*/
func CalculatePrime() {
	var nThInteger int
	fmt.Println("Program to print all the prime numbers up to the given nTh integer. Including the nTh integer")
	fmt.Println("enter the value of nth integer")
	fmt.Scanln(&nThInteger)
	printPrimeTill(nThInteger)
}

// The function `printPrimeTill` prints all prime numbers up to a given integer.
func printPrimeTill(nThInteger int) {
	fmt.Println("Primes up to ", nThInteger)
	for currentInt := 2; currentInt <= nThInteger; currentInt++ {
		if currentInt == 2 {
			fmt.Printf("%d \t", currentInt)
			continue
		} else if currentInt%2 == 0 {
			continue
		} else if checkPrime(currentInt) {
			fmt.Printf("%d \t", currentInt)
		}
	}
	fmt.Println()
}

/*
The function "checkPrime" checks if a given integer is prime or not.
The function checks if a given integer is prime by iterating odd numbers from 3 to the square root of the
integer and checking for divisibility.
*/
func checkPrime(currentInt int) bool {
	sqrtCurrentInt := int(math.Sqrt(float64(currentInt)))
	for i := 2; i <= sqrtCurrentInt; i += 1 {
		if currentInt%i == 0 {
			return false
		}
	}
	return true
}

package pn14

import (
	"fmt"
)

/*
Write a Go function to print the reverse of the string and check for palindrome
for example
input:""
output:""
*/

func ReverseTheString() {
	fmt.Println("enter the string to reverse it and check for palindrome")
	var input string
	fmt.Scanln(&input)
	reverseString := getReverse([]byte(input))
	fmt.Printf("the %v reverse is %v \n", input, reverseString)
	if checkPalindrome(input, reverseString) {
		fmt.Println("and yes it is palindrome")
	} else {
		fmt.Println("and it is not a palindrome")
	}

}

func getReverse(inputBytes []byte) string {
	i := 0
	j := len(inputBytes) - 1
	for i < j {
		inputBytes[i], inputBytes[j] = inputBytes[j], inputBytes[i]
		i += 1
		j -= 1
	}
	return string(inputBytes)
}

func checkPalindrome(input string, reverseString string) bool {
	if input == reverseString {
		return true
	}
	return false
}

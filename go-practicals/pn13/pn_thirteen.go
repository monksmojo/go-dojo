package pn13

import (
	"fmt"
	"strings"
)

/*
Write a Go function to remove duplicate characters of a given string.
for example
input:""
output:""
*/

func RemoveDuplicates() {
	fmt.Println("enter a string to remove the duplicate charters from it")
	var inputString string
	fmt.Scanln(&inputString)
	outputString := getUniqueCharacters(inputString)
	fmt.Printf("the %v string after removing duplicate character is %v \n", inputString, outputString)
}

func getUniqueCharacters(inputString string) string {
	var uniqueChar strings.Builder
	charCheckMap := make(map[rune]bool)
	for _, char := range inputString {
		_, ok := charCheckMap[char]
		if !ok {
			uniqueChar.WriteRune(char)
			charCheckMap[char] = true
		}
	}
	return uniqueChar.String()
}

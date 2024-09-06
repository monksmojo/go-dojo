package pn12

import (
	"fmt"
	"strings"
	"unicode"
)

/*
PN12: Write a Go function to convert a given string to all uppercase if it contains at least uppercase characters in the first 4 characters.
input: Chocolate
output: CHOCOLATE
input: happyBirthDay
output: happyBirthDay
*/

func StrToUpper() {
	fmt.Println("1. select version 1 to execute the problem statement")
	fmt.Println("2. select version 2 to execute the problem statement")
	var choice int
	var output string
	fmt.Println("Choose the version")
	fmt.Scanln(&choice)
	inputString := getInput()
	switch choice {
	case 1:
		output = toUpperV1(inputString)

	case 2:
		output = toUpperV2(inputString)
	}
	fmt.Println(output)
}

func getInput() (inputString string) {
	fmt.Println("enter a string to check weather it can be converted to uppercase or not")
	fmt.Scanln(&inputString)
	return inputString
}

func toUpperV1(stringArg string) string {
	for i := 0; i < 4; i++ {
		if unicode.IsUpper(rune(stringArg[i])) {
			fmt.Printf("%v index character %v found in uppercase \n", i, stringArg[i])
			return fmt.Sprintln(stringArg, " in uppercase will be ", strings.ToUpper(stringArg))
		}
	}
	return "no upperCase character found in first four character of the string"
}

func toUpperV2(stringArgs string) string {
	for i, v := range stringArgs {
		if i <= 3 && unicode.IsUpper(v) {
			fmt.Printf("%v index character found in uppercase \n", i)
			return fmt.Sprintln(stringArgs, " in uppercase will be ", strings.ToUpper(stringArgs))
		}
	}
	return "no upperCase character found in first four character of the string"
}

package pn7

import "fmt"

/*
Write a Go function to swap the consecutive characters in a string input by the user.
For example
input: “chocolate”
output: "hccolotae"
*/

func SwapCharactersOfString() {
	fmt.Println("Choose the method to swap the characters")
	fmt.Println("1.) By appending swap characters to the new byte slice and returning byte slice type casted as string")
	fmt.Println("2.) By converting string into rue slice swapping characters and returning rune slice type casted as string")
	var choice int
	fmt.Scanln(&choice)
	fmt.Println("Give me a string to swap the consecutive characters : ")
	var inputString string
	fmt.Scanln(&inputString)
	switch choice {
	case 1:
		swapCharacterV1(inputString)
	case 2:
		swapCharacterV2(inputString)
	}
	outputString := swapCharacterV1(inputString)
	fmt.Println("the given input string is ", inputString)
	fmt.Println("the output swapped character string is ", outputString)

}

func swapCharacterV1(str string) string {
	strLength := len(str)
	var swapBytes = make([]byte, 0, strLength)
	fmt.Println()
	for i := 0; i < strLength-1; i = i + 2 {
		swapBytes = append(swapBytes, str[i+1], str[i])
	}
	if strLength%2 != 0 {
		swapBytes = append(swapBytes, str[strLength-1])
	}
	return string(swapBytes)
}

func swapCharacterV2(inputString string) string {
	runeSlice := []rune(inputString)
	for i := 0; i < len(runeSlice)-1; i += 2 {
		runeSlice[i], runeSlice[i+1] = runeSlice[i+1], runeSlice[i]
	}
	return string(runeSlice)
}

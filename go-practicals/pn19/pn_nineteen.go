package pn19

import (
	"fmt"
	"strconv"
)

func SearchCharacterSlice() {

	str := "OLYMPIA"
	character := "A"
	var index int
	var subChoice int
	fmt.Println("we have two methods to solve this practical")
	fmt.Println("1. using range to iterate and search")
	fmt.Println("2. using index to iterate and search")
	fmt.Println("enter the choice to solve the practical")
	fmt.Scanln(&subChoice)
	switch subChoice {
	case 1:
		fmt.Print("iterate using range ")
		index = searchUsingRange(str, character)
	case 2:
		fmt.Println("iterate using index")
		index = searchUsingIndex(str, character)
	default:
		fmt.Println("invalid choice")
	}

	if index != -1 {
		fmt.Printf("found character %s in string %s at position %d \n", character, str, index+1)
	} else {
		fmt.Printf("%s character not found in string %s at position %d \n", character, str, index+1)
	}
}

func searchUsingRange(str, character string) int {
	target, err := strconv.ParseUint(character, 10, 32)
	if err != nil {
		fmt.Println("error converting string to char")
		return -1
	}
	for i, v := range str {
		if v == rune(target) {
			return i
		}
	}
	return -1
}

func searchUsingIndex(str, character string) int {
	target, err := strconv.ParseUint(character, 10, 32)
	if err != nil {
		fmt.Println("error converting string to char")
		return -1
	}
	for i := 0; i < len(str); i++ {
		if str[i] == byte(target) {
			return i
		}
	}
	return -1
}

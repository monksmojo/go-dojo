package pn9

import "fmt"

/*
Write a Python program to remove the characters which have odd index values of a given
string
for ex: 
input: "pokemon"
output: "pkmn"
*/

func FilterOddIndex() {
	var inputString string
	fmt.Println("Enter the string you want get filtered")
	fmt.Scanln(&inputString)
	r:=getOddFilteredSlice(inputString)
	fmt.Println("original string",inputString)
	fmt.Println("string after odd index removed",r)
}

func getOddFilteredSlice(givenString string) string {
	evenCharacterSlice := []rune{}
	for index, value := range givenString {
		if index%2 == 0 {
			evenCharacterSlice = append(evenCharacterSlice, value)
		}
	}
	return string(evenCharacterSlice)
}

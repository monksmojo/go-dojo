package pn15

import (
	"fmt"
	"strings"
)

/*
PN:15 Write a Go function to get a single string from two given strings, separated by a space and swap the first two characters of each string. Eg input: 'abc', 'xyz' output: 'xyc abz'
*/

func CreateString() {
	fmt.Println("Create single string from two given strings")
	var inputString1 string
	var inputString2 string
	fmt.Println("enter the 1st string ")
	fmt.Scanln(&inputString1)
	fmt.Println("enter the 2nd string ")
	fmt.Scanln(&inputString2)
	r := getSingleString(inputString1, inputString2)
	fmt.Println("1 single string generated from the input 2 string is ", r)
}

func getSingleString(str1, str2 string) string {
	var singleStr strings.Builder
	singleStr.WriteByte(str2[0])
	singleStr.WriteString(str1[1:])
	singleStr.WriteRune(' ')
	singleStr.WriteByte(str1[0])
	singleStr.WriteString(str2[1:])
	return singleStr.String()
}

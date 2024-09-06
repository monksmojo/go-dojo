package pn8

import (
	"fmt"
	"strings"
)

/*
Write a Python program to input 2 strings i.e. s1 and s2 and find the presence of s1 in s2,
if present then replace it with equal number of # in it. Eg input: s1=”hi”
s2=”hello_hi_welcome” , output: hello_##_welcome
*/

func EncodeTheString() {
	var parentString string
	var childString string
	fmt.Println("Enter the parent string:")
	fmt.Scanln(&parentString)
	fmt.Println("Enter the child string:")
	fmt.Scanln(&childString)
	fmt.Println("Encoding the string if child string is present in the parent string")
	r := getEncodeString(childString, parentString)
	fmt.Println(r)
}

func getEncodeString(childString string, parentString string) string {
	hashString := strings.Repeat("#", len(childString))
	if strings.Contains(parentString, childString) {
		encodedString := strings.ReplaceAll(parentString, childString, hashString)
		return fmt.Sprintf("encoded string is %s \n", encodedString)
	}
	return fmt.Sprintf("%s does not exist in %s \n", childString, parentString)
}

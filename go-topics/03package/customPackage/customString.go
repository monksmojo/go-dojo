package customPackage

import (
	"fmt"
	"strings"
)

// come to this parts of notes when you hae completed the
//

var StrPackageName = "this is a custom string package"

/**
Compare()	checks if two strings are equal
Contains()	checks if the string contains a substring
Count()	counts the number of times a substring present in the string
Join()	creates a new string by concatenating elements of a string array
ToLower()	converts the string to lowercase
ToUpper()	converts the string to uppercase
**/

var str1 string = "clock"
var str2 string = "tower"
var str3 string = "red color red challenger "
var slice1 = []string{"q", "w", "e", "r", "t", "y"}

func GetStr1() string {
	return str1
}

func GetStr2() string {
	return str2
}

func GetStr3() string {
	return str3
}

func GetSlice() []string {
	return slice1
}

func Compare2Str(str1 string, str2 string) string {
	i := strings.Compare(str1, str2)
	if i == 0 {
		return fmt.Sprintf("%s both strings are equal \n", str1)
	} else if i > 0 {
		return fmt.Sprintf("%v is greater string \n", str1)
	} else {
		return fmt.Sprintf("%v is greater string \n", str2)
	}
}

func IfContains(parentString string, childString string) string {
	if strings.Contains(parentString, childString) {
		return fmt.Sprintf("%v contains %v \n", parentString, childString)
	} else {
		return fmt.Sprintf("%v does not contain %v \n", 	parentString, childString)
	}
}

func CountTheOccurrence(stringToCount string, parentString string) int {
	return strings.Count(parentString, stringToCount)
}

func ArrayToString(array []string) string {
	return strings.Join(array, "")
}

func StringInCase(str string, caseSelected string) string {
	switch strings.ToLower(caseSelected) {
	case "upper":
		return strings.ToUpper(str)
	case "lower":
		return strings.ToLower(str)
	default:
		return fmt.Sprintf("%v no such case is present to convert %v \n", caseSelected, str)
	}
}

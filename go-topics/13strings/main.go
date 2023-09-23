package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to the Strings in go")
	// declaring String
	var myStr string
	fmt.Printf("the default value of a string is %v \n", myStr)

	// declare and initialize the string method-1
	var yourStr string = "popsicle"
	fmt.Printf("string by method-1 %v \n", yourStr)

	// declare and initialize the string method-2
	var popStr = "popsicle"
	fmt.Printf("string by method-2 %v \n", popStr)

	// declare and initialize the string method-3
	iceStr := "iceCream"
	fmt.Printf("string by method-3 %v \n", iceStr)

	// declare and initialize the string method-4
	// we can use back ticks to represent strings
	var msgTime string = `welcome to the endless time`
	fmt.Printf("string by method-4 %v \n", msgTime)
	// that backticks can be used to declare multi-line strings or strings that include special characters without having to escape them.
	// difference between `` and "" while declaring string
	// "" double quotes it treats the special/escape character in the string
	var escapeStr string = "Hey ! \t HI \n Welcome"
	fmt.Println(escapeStr)
	// `` back ticks it escapes the special character and escape characters in string
	var noEscapeStr string = `Hey ! \t HI \n Welcome`
	fmt.Println(noEscapeStr)

	// accessing single characters of the string
	// indexing

	var name1 string = "Collins"
	// 1st character  of the string

	fmt.Printf("1st character of %s is %c \n", name1, name1[0])
	fmt.Printf("last character of %s is %c \n", name1, name1[6])
	fmt.Printf("middle character of %s is %c \n", name1, name1[4])

	// finding the length of the string
	// we can find the number iof characters in the string
	// using len() built in function
	// len() returns int containing
	// the number of characters in the function-

	numberOfCharacters := len(name1)
	fmt.Printf("the number of characters in %s is %d \n", name1, numberOfCharacters)

	// concatenation of two strings
	// joining of two strings
	// we use '+' operator to join two strings
	var msg1 string = "the prefix"
	var msg2 string = "the suffix"
	msg3 := msg1 + msg2
	fmt.Printf("concatenation of two strings %s + %s into a new string %s \n", msg1, msg2, msg3)

	// some important golang string methods
	//  1. compare() compare two strings
	// comparing is done lexicographically and is case sensitive
	var name2 string = "app"
	var name3 string = "zed"
	var name4 string = "APP"
	i1 := strings.Compare(name2, name3)
	fmt.Println("i1", i1)
	//  -1 because name2 <(less than) name3
	i2 := strings.Compare(name2, name1) // name1 is Collins
	fmt.Println("i2", i2)
	// 1 because name2 >(greater than) name3
	i3 := strings.Compare(name2, name4)
	fmt.Println("i3", i3)
	// 1 both are equal but case sensitive
	// ascii value of "a" is >(greater than) ascii values of "A"

	// 2. contains("parentString","substring")
	// the strings.Contains("parentString","substring") tells weather ch substring is part of parentString or not returns Boolean
	var ch string = "A"
	var vowels string = "aeiouAEIOU"
	isVowel := strings.Contains(vowels, ch)
	if isVowel {
		fmt.Println(ch, "is vowel")
	} else {
		fmt.Println(ch, "is not vowel")
	}

	// 3. replace("parentStr","oldString","newString","numberOfTimes")
	// strings.Replace("parentStr","oldString","newString","numberOfTimes")
	// it is used to replace old string with new string
	var thing string = "copper bottle"
	newThing := strings.Replace(thing, "bottle", "wire", 1)
	fmt.Println("old string", thing, "is updated using replace() function", newThing)

	// 4. toUpper("string")
	// newString:=strings.ToUpper(string)
	// converts the string argument to upperCase and returns a newString
	var carLowerStr string = "grand"
	var carUpperStr string = strings.ToUpper(carLowerStr)
	fmt.Printf("converting lowercase %s into uppercase %s \n", carLowerStr, carUpperStr)

	// 5. toLower("string")
	// newString:=strings.ToLower(string)
	// converts the string argument to upperCase and returns a newString
	carLowerStr = strings.ToLower(carUpperStr)
	fmt.Printf("converting the uppercase %v into lowercase %v \n", carUpperStr, carLowerStr)

	// 6. split()
	// newString:= strings.Split("string","delimiter")
	// use to split a string into multiple sub string based on the delimiter
	var email string = "mikeySan@gmail.com"
	name5 := strings.Split(email, "@")
	fmt.Printf("splitting email %v with @ to get %v of type %T \n", email, name5, name5)

	// type conversion of string data type to another data type

	// string character to int ascii value
	cAsciiValue := rune('C')
	aAsciiValue := int('a')
	fmt.Println("ascii value of C", cAsciiValue)
	fmt.Println("ascii value of a", aAsciiValue)

	// string digit to int
	var inputValue string = "67"
	value, err := strconv.Atoi(inputValue)
	if err != nil {
		fmt.Println("the input value is not a digit string")
	}
	fmt.Printf("str value %v of type %T converted to %v of type %T \n", inputValue, inputValue, value, value)

	var yoMiked string = "yo miked"

	// Convert the string to a []rune character array.
	// In Golang, a character is represented by a rune, which is a 32-bit integer.
	// runes allow Golang to represent any character in any encoding
	runeCharArray := []rune(yoMiked)
	fmt.Printf("converting the string %v into %v rune array \n", yoMiked, runeCharArray)

	// Convert the string to a []byte character array
	// where each byte represents a single character in the string.
	byteCharArray := []byte(yoMiked)
	fmt.Printf("converting the string %v into %v byte array \n", yoMiked, byteCharArray)

	// The main difference between []rune and []byte is the type of data they represent.
	// a 'rune' which is an alias for int32.
	// It is commonly used when working with Unicode characters or strings that may contain non-ASCII characters.
	// while []byte is used for working with byte data.
	//  which is an alias for uint8. It is commonly used when dealing with binary data, raw bytes, or ASCII characters. It is also used when working with file I/O or network communication.

	// Escape Sequence in Golang String
	// some times we want to display '' or "" inside string
	// for ex  "he said ! watch out"
	// we can use `` back ticks to create a string we ignore escape sequence
	var msg4 string = `if you "creeping" don't let me show`
	fmt.Println(msg4)

	// in go strings are immutable

	// 	// string formatting
	// 	### Format specifiers
	// - Go language provides various format specifiers for different types of variables. Here are some commonly used format specifiers in Go:

	// - %d - decimal integer
	// - %f - floating-point number
	// - %s - string
	// - %t - boolean
	// - %v - any value
	// - %T - type of the variable
	// - %p - pointer

	// go lang provides fmt.Sprintf() function for string formatting
	msg := fmt.Sprintf("hi my name is %s and i am %d years old", "Fuji", 22)
	fmt.Println(msg)
}

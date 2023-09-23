package main

import "fmt"

// global public constant
const globalPublicConstToken string = "qwerty"

func main() {
	// String
	var username string = "hey monkey"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)
	// %T is used to give type of the variable

	// Boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type %T \n", isLoggedIn)

	// unsigned integer
	// it cannot store negative values
	var smallInt8 uint8 = 255                    // maximum value of uint8
	var smallInt16 uint16 = 65535                // maximum value of uint16
	var smallInt32 uint32 = 4294967295           // maximum value of uint32
	var smallInt64 uint64 = 18446744073709551615 // maximum value of uint64
	var uInteger uint = 898
	fmt.Println(smallInt8)
	fmt.Printf("variable is of type %T \n", smallInt8)
	fmt.Println(smallInt16)
	fmt.Printf("variable is of type %T \n", smallInt16)
	fmt.Println(smallInt32)
	fmt.Printf("Variable is of type %T \n", smallInt32)
	fmt.Println(smallInt64)
	fmt.Printf("Variable is of type %T \n", smallInt64)
	fmt.Println(uInteger)
	fmt.Printf("Variable is of type %T \n", uInteger)

	// signed integer
	// it can store negative values
	var sInteger int = 34565
	fmt.Println(sInteger)
	fmt.Printf("Variable is of type %T \n", sInteger)

	// by default signed integer and unsigned integer have a alias int and uint
	// size of int and uint depend on the architecture of the machine
	// if it is 64 bit machine it will work as int64 and uint64

	// float32 is a 32 bit float that gives values till 5 digits after decimal point
	var smallFloat float32 = 4786.890567345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of Type %T\n", smallFloat)

	// float64 is a 64 bit float that gives more precision then float32
	var float float64 = 4786.890567345
	fmt.Println(float)
	fmt.Printf("Variable is of Type %T\n", float)

	// rune used for unicode characters , internally used as 32-bit integer
	var character rune = 'a'
	fmt.Println(character)
	fmt.Printf("character %c is of type %T \n", character, character)

	// default values and alias
	// default value of int
	var defaultInt int
	fmt.Printf("default value of int is %v \n", defaultInt)
	// default value of float
	var defaultFloat float64
	fmt.Printf("default value of float is %v \n", defaultFloat)
	// default value of string
	var defaultString string
	fmt.Printf("default value of string is %v \n", defaultString)
	// default value of boolean
	var defaultBool bool
	fmt.Printf("default value of bool id %v \n", defaultBool)

	// implicit type declaration and initialization
	var autoVariable = 234
	fmt.Println(autoVariable)
	fmt.Printf("type of the implicit variable is %T \n", autoVariable)

	// implicit variable declaration using walrus operator (no var way)
	autoVariable2 := "fist of fury"
	fmt.Println(autoVariable2)
	fmt.Printf("type of the implicit variable2 is %T \n", autoVariable2)
	// No, once a variable is declared with an data type in Go, its data type is fixed and cannot be changed.

	// we cannot declare variable with := (walrus operator) outside function in go
	var iCanBeGlobal int = 101
	fmt.Printf(" i can be a global variable %v \n", iCanBeGlobal)

	// accessing global public variable
	fmt.Println(globalPublicConstToken)
	fmt.Printf("type of global public constant globalPublicConstToken is %T \n", globalPublicConstToken)

	// creating multiple elements in go method 1
	var name, age = "camel", 90
	fmt.Printf("my name is %v and age is %v \n", name, age)

	// creating multiple elements in go using walrus operator
	myName, myAge := "camel", 90
	fmt.Printf("my name is %v and age is %v \n", myName, myAge)

	// rune // it is used to
	var ch rune = 'Z'
	fmt.Println("character", ch)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome type casting in go language")
	// Type casting is the process of converting the value of one data type

	// for example to convert
	var intNumber int = 32
	// int to float32
	var intToFloat float32 = float32(intNumber)
	fmt.Printf("int value %d to %f value \n", intNumber, intToFloat)
	fmt.Printf("value %f data type is of type %T \n", intToFloat, intToFloat)
	// int to string
	var intToString string = strconv.Itoa(intNumber)
	fmt.Printf("int value %d to %s value \n", intNumber, intToString)
	fmt.Printf("value %s data type is of type %T \n", intToString, intToString)
	// int to bool
	// there is no direct way to convert int to boolean in go

	var floatNumber float32 = 32.00
	// float32 to int
	var floatToInt int = int(floatNumber)
	fmt.Printf("float value %f to %d int value \n", floatNumber, floatToInt)
	fmt.Printf("value %d data type is of type %T \n", floatToInt, floatToInt)
	// float32 to string
	var floatToString string = strconv.FormatFloat(float64(floatNumber), 'f', 2, 32)
	// it takes 4 value
	// strconv.FormatFloat(float64(value),'f',precision,bit_size)
	// 'f' is format specifier
	//  in our case 32 bit_size
	fmt.Printf("float value %f to %s value \n", floatNumber, floatToString)
	fmt.Printf("value %s data type is of type %T \n", floatToString, floatToString)
	// int to bool
	// there is no direct way to convert int to boolean in go

	// string to integer
	// string to integer method-1
	// method 1 converts string to int
	var str string = "10"
	stringToInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("error converting string to int")
	} else {
		fmt.Printf("%v converted to %v having type %T \n", str, stringToInt, stringToInt)
	}

	// string to integer
	// string to integer method-2
	// method 2 converts string to int64 or int32
	stringToInt64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("error converting string to int64")
	} else {
		fmt.Printf("%v converted to %v having type %T \n", str, stringToInt64, stringToInt64)
	}

	// string to float32 or float64
	var str2 string = "12.43"
	stringToFloat32, err := strconv.ParseFloat(str2, 32)
	if err != nil {
		fmt.Println("error converting string to float32")
	} else {
		fmt.Printf("%v converted to %v having type %T \n", str2, stringToFloat32, stringToFloat32)
	}

	// string to slice
	var st1 string = "copper mesh"
	st1Slice := []rune(st1)
	fmt.Printf("converting string %v into %v \n", st1, st1Slice)

	// converting a slice of []rune{} characters into the string
	for _, v := range st1Slice {
		fmt.Printf("%v : %v\n", v, string(v))
	}
	fmt.Println()

	// converting a slice of []rune{} characters into a single string
	fmt.Println(string(st1))

	// taking input using bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to our pizza parlour")
	fmt.Println("please rate our pizza from 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating: ", userInput)
	fmt.Printf("rating : is of Type %T \n", userInput)

	// type converting above string ito float for further conversion
	rating, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("the value enter by the user is not of proper format")
	} else {
		var ratingPercent float64 = (rating / 5.0) * 100
		var ratingOutOf10 float64 = (ratingPercent / 100) * 10.0
		fmt.Println("the ratings we get out of 10 are", ratingOutOf10)
	}

	// taking input using bufio.NewScanner(os.Stdin)
	fmt.Println("welcome to bookCab")
	fmt.Println("please rate our application from 1 to 5")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput = scanner.Text()
	fmt.Println("Thanks for the rating: ", userInput)
	fmt.Printf("rating : is of Type %T \n", userInput)
	rating, err = strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("the value enter by the user is not of proper format")
	} else {
		var ratingPercent float64 = (rating / 5.0) * 100
		var ratingOutOf10 float64 = (ratingPercent / 100) * 10.0
		fmt.Println("the ratings we get out of 10 are", ratingOutOf10)
	}

}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")
	fmt.Println("array are fixed size container")

	// declaring an array

	var intArray [4]int
	fmt.Println("default value of int array is", intArray)
	fmt.Printf("type of the intArray is %T \n", intArray)
	var floatArray [4]float64
	fmt.Println("default value of float array is", floatArray)
	var stringArray [4]string
	fmt.Println("default value of String array is", stringArray)
	fmt.Printf("type of the stringArray is %T \n", stringArray)
	var boolArray [4]bool
	fmt.Println("default value of bool array is", boolArray)

	// array are fixed size container
	// it is a collection of similar data type values
	// default value of int array is [0 0 0 0]
	// default value of float array is [0 0 0 0]
	// default value of String array is [   ]
	// default value of bool array is [false false false false]

	//   declaring and initializing the array method 1
	var fruitArray [5]string
	fruitArray[0] = "apple"
	fruitArray[1] = "mango"
	fruitArray[2] = "peach"
	fruitArray[3] = "banana"
	fmt.Println(fruitArray)

	//   declaring and initializing the array method 2
	// syntax
	//  var array_variable = [size]datatype{elements of array}
	var country = [5]string{"India", "Poland", "Malaysia", "Tibet"}
	fmt.Println(country)

	// the array are mutable for the values
	// we can add value as long as it does not exceed its declared size
	//  we can update the values of the array but not the size
	country[4] = "vietnam"
	fmt.Println("adding to the elements of the array", country)

	country[3] = "New Zealand"
	fmt.Println("updating the elements of the array", country)
	fmt.Printf("type of the country is %T \n", country)

	// method 3
	// declaring and initializing the array
	// here the Go lang compiler infer it from the number of elements provided in the array literals
	var appliances = [...]string{"mouse", "keyboard", "screen", "mic"}
	fmt.Println("appliances array made at run time", appliances)
	fmt.Printf("type of the appliances is %T \n", appliances)

	// method 4
	// declaring and initializing the array with specific index value
	// i.e Initialize specific elements of an array
	var arrayOfStr = [5]string{0: "p", 1: "q", 4: "r"}
	fmt.Println("value in arrayOfStr is", arrayOfStr)
	// since default value of string array is '' (clank string)
	// so for index 2,3 we have blank string '' as value

	// further changing the value at index 3
	arrayOfStr[3] = "b"
	fmt.Println("updated value in arrayOfStr is", arrayOfStr)

	// method 5
	// creating a multidimensional array
	var twoDMatrix = [2][2]string{{"q", "w"}, {"e", "r"}}
	fmt.Println("the multidimensional array value", twoDMatrix)
	twoDMatrix[0][0] = "Q"
	fmt.Println("updated the multidimensional array value", twoDMatrix)

	// size of the array is fixed at the time of declaration and does not depend on the number of elements the array will hold
	var things = [5]string{"keyChain", "apple", "cup"}
	// we use len() built in function to get the length of the array
	size := len(things)
	fmt.Println("size of array", size) // size of array 5
	// size would be 5 even if we have only 3 elements in the array

	// in go the array are passed as value to the function
	// i.e a copy of the array is created and sent to the function
	var food = [3]string{"pasta", "pizza", "noodles"}
	fmt.Printf("food array before the function call %v address %p \n", food, &food)
	updateFoodList(food)
	fmt.Printf("food array after the function call %v address %p \n", food, &food)

}

func updateFoodList(foodList [3]string) {
	foodList[2] = "dumplings"
	fmt.Println("updating food array inside the function")
	fmt.Printf("food array inside the function %v address %p \n", foodList, &foodList)
}

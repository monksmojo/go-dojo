package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("welcome to the slice in go")
	// slice are collection of similar data types like array
	// however unlike array slice don't have fixed size
	// we can add or remove element from the slice

	// different ways of creating the slice
	// method -1 declaring an empty slice
	var numbers []int
	fmt.Println("a empty integer slice", numbers)
	fmt.Printf("type of this slice is %T \n", numbers)

	//method-2 declaring and initializing slice
	var intNumber = []int{21, 22, 24, 25, 32}
	fmt.Println("a integer slice", intNumber)
	fmt.Printf("type of this slice is %T \n", intNumber)

	// method-3
	// using  () method to create slice
	// make(type,size_of_slice,max_capacity_of_slice(optional))
	name := make([]string, 3, 7)
	// 3 is the length of the slice (number of elements present in the slice)
	// 7 is the capacity of the slice (maximum size up to which a slice can be extended)
	name[0] = "rich"
	name[1] = "jill seth"
	name[2] = "mike rico"
	fmt.Println("slice using make function", name)

	// if you notice above a slice don't have size defined
	// while declaration [] is empty

	// declaring an array
	var intArray = [4]int{20, 38, 67, 89}
	// if you notice above a array has defined size
	// while declaration [4] is empty
	fmt.Println("a integer slice", intArray)
	fmt.Printf("type of this slice is %T \n", intArray)

	// creating a slice from array
	charArray := [5]string{"q", "w", "e", "r", "t"}
	var charSlice []string = charArray[0:5]
	// while creating a slice from an array we use [start:stop+1]
	fmt.Println("the charArray is ", charArray)
	fmt.Println("char slice created from array is ", charSlice)

	// creating a slice from a string
	var charString string = "violet"
	var strSlice = charString[3:] // here the data type of the strSlice will be decided at the run time
	fmt.Println("the string of char is", charString)
	fmt.Println("string slice created from string is ", strSlice)

	// function on the slice
	// append()
	// we use append() function to add elements to the slice
	primeNumbers := []int{2, 3, 5}
	fmt.Println("prime number slice", primeNumbers)
	primeNumbers = append(primeNumbers, 9, 11)
	fmt.Println("appending more elements to the the primeNumber slice", primeNumbers)

	// appending one slice at the end of another slice
	var myFruits = []string{"apple", "banana"}
	var newFruits = []string{"orange", "kiwi"}
	fmt.Printf("adding new fruits %v to my fruits %v \n", newFruits, myFruits)
	myFruits = append(myFruits, newFruits...)
	fmt.Printf("fruits i have %v \n", myFruits)

	// appending two slice together to create new slice
	evenNumbers := []int{2, 4, 6, 8}
	fmt.Println("concatenation ")
	oddNumbers := []int{1, 3, 5, 7}
	var fewNumbers []int = append(evenNumbers, oddNumbers...)
	fmt.Println("slice of even number", evenNumbers)
	fmt.Println("slice of odd number", oddNumbers)
	fmt.Printf("adding two slice %v and %v into one %v \n", evenNumbers, oddNumbers, fewNumbers)

	// copy()
	// copy(destination,source)
	// copy the source element "onto" the destination
	// here we have put emphasis on "onto" as
	// the source slice elements overlaps on the destination elements
	// the number of elements overlaps depends upon the size of the destination
	var naturalNumbers = []int{9, 10}
	fmt.Printf("overlapping natural  numbers %v with even numbers %v to create new natural numbers %v using copy() \n", naturalNumbers, evenNumbers, copy(naturalNumbers, evenNumbers))
	// copy(naturalNumbers, evenNumbers
	// reflect.DeepEqual()
	// comparing of two slice
	// reflect.DeepEqual() function to compare these slice. The function returns true if both the slice have the same length and same elements in the same order, otherwise it returns false.
	a := reflect.DeepEqual(naturalNumbers, evenNumbers)
	fmt.Printf("does %v equal %v %t \n", naturalNumbers, evenNumbers, a)
	// len()
	// it is used to get number of elements in the slice
	var manyThings = []string{"tab", "smartphone", "laptop"}
	fmt.Printf("the number of elements in the slice %v are %v \n", manyThings, len(manyThings))

	// accessing the elements of the slice
	var item1 string = manyThings[0]
	fmt.Println("fetched 0 index value ", item1)

	// updating the element of the slice
	manyThings[0] = "tablet"
	fmt.Println("updated slice", manyThings)

	// in go the slices are passed as reference to the function
	// i.e a reference to the memory address of the slice is sent to the function
	var food = []string{"pasta", "pizza", "noodles"}
	fmt.Printf("food slice before the function call %v address %p \n", food, &food)
	updateFoodList(food)
	fmt.Printf("food slice after the function call %v address %p \n", food, &food)

	// when passing n number of arguments to the function
	// the function becomes the variadic function and receives variadic arguments
	s := sumOfIntegers(29, 26, 25, 24)
	fmt.Println("sum of n integers is ", s)

}

// The function `updateFoodList` updates the third element of a given slice and prints the updated list.
func updateFoodList(foodList []string) {
	foodList[2] = "dumplings"
	fmt.Println("updating food array inside the function")
	fmt.Printf("food slice inside the function %v address %p \n", foodList, &foodList)
}

// the variadic function receives variadic arguments as a slice
func sumOfIntegers(intList ...int) (sum int) {
	fmt.Printf("the intList is %v and its data type is %T \n", intList, intList)
	for _, element := range intList {
		sum += element
	}
	return sum
}

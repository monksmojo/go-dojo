package pn2

import (
	"fmt"
	"sort"
)

// Practical2 Description
/**
@initialAuthor: monks_mojo
@lastUpdatedBY: monks_mojo
Practical Two:
WAF which takes an array as an argument and reverse the array


Test case sample:1
input: given the array : [2,5,7,3,17,13,11]
output: reversed array : [11,13,17,3,7,5,2]
**/

/**
we are gonna use three ways to reverse the array
1. reverse the array using the naive approach
2. reverse the array using two pointers approach
3. Reversing the slice in descending order using the built in function
**/

// The function ReverseArray allows the user to choose between three different approaches to reverse an array.
func ReverseArray() {
	fmt.Println("1. reverse the array using the naive approach.")
	fmt.Println("2. reverse the array using two pointers approach.")
	fmt.Println("3. reverse the array using built-in functions.")
	fmt.Println("enter your choice:")
	var ch int
	fmt.Scanln(&ch)
	var intArray = [7]int{2, 5, 7, 3, 17, 13, 11}
	switch ch {
	case 1:
		fmt.Println("reversing the array using the naive approach.")
		reverseArrayType1(intArray)
	case 2:
		fmt.Println("reversing the array using the two pointer approach.")
		reverseArrayType2(intArray)
	case 3:
		fmt.Println("reversing the slice in descending order using the built in function")
		reverseArrayType3(intArray[:])
	default:
		fmt.Println("you entered the wrong choice dummy !")
	}
}

// The function reverseArrayType1 takes an array of integers as input and reverses the order of its elements.
// Using the naive approach
func reverseArrayType1(arr [7]int) {
	fmt.Printf("original array %v memory location %p \n", arr, &arr)
	size := len(arr)
	for i := 0; i < size/2; i += 1 {
		c := arr[i]
		arr[i] = arr[size-i-1]
		arr[size-i-1] = c
	}
	fmt.Printf("reversed array %v memory location %p \n", arr, &arr)
}

// The function reverseArrayType2 takes an array of integers as input and reverses the order of its elements.
// Using the two pointers approach
func reverseArrayType2(arr [7]int) {
	fmt.Printf("original array %v memory location %p \n", arr, &arr)
	size := len(arr)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Printf("reversed array %v memory location %p \n", arr, &arr)
}

// The function takes an integer slice, sorts it in reverse order, and prints the original and sorted
// slices along with their memory locations.
func reverseArrayType3(arrSlice []int) {
	fmt.Printf("original slice %v memory location %p \n", arrSlice, &arrSlice)
	reverseArraySlice := sort.Reverse(sort.IntSlice(arrSlice))
	sort.Sort(reverseArraySlice)
	fmt.Printf("sortReversed slice %v memory location %p \n", arrSlice, &arrSlice)
}

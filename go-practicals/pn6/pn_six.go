package pn6

import "fmt"

/*
Write a Go function `filterSlice` that filter the multiples of 3 and 5 from the slice and returns the remaining slice
Input:
[]int{13,11,3,6,9,15,22}
Output:
[]int{13,11,22}

We will use two methods:
1. create a new slice to store all non multiples of 3 and 5
2. reference the original slice that only contain the non multiple of 3 and 5
*/


func FilterSlice() {
	var numberSlice = []int{13, 11, 3, 6, 9, 15, 22}
	fmt.Printf("%v Filter multiple of 3 and 5 from the slice \n", numberSlice)
	fmt.Println("1. To Filter the multiples by creating the new slice")
	fmt.Println("2. To Filter the multiples by referencing the original slice")
	var choice int
	fmt.Println("enter your choice: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		filterByNewSlice(numberSlice)
	case 2:
		filterBySliceReference(numberSlice)
	default:
		fmt.Println("Not the correct choice dummy")
	}
}


func filterByNewSlice(numberSlice []int) {
	var filterSlice = []int{}
	for _, v := range numberSlice {
		if v%3 != 0 && v%5 != 0 {
			filterSlice = append(filterSlice, v)
		}
	}
	fmt.Printf("slice with non multiple of 3 and 5 is %v \n", filterSlice)
}


func filterBySliceReference(numberSlice []int) {
	i := len(numberSlice) - 1
	for j := 0; j < i; j++ {
		if numberSlice[j]%3 == 0 || numberSlice[j]%5 == 0 {
			numberSlice[i], numberSlice[j] = numberSlice[j], numberSlice[i]
			j -= 1
			i -= 1
		}
	}
	fmt.Printf("slice with non multiple of 3 and 5 is %v \n", numberSlice[:i])
}

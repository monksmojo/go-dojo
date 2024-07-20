package pn18

import "fmt"

// PN:18LinearSearch: search an int in slice of int
// Given a target is integer ex 20
// Given a slice
func SearchInIntSlice() {
	intSlice := []int{10, 20, 13, 14, 18, 12}
	target := 14
	var subChoice int
	fmt.Println("we have two methods to solve this practical")
	fmt.Println("1. search the element linearly")
	fmt.Println("2. search using two pointers approach")
	fmt.Println("enter your choice")
	fmt.Scanln(&subChoice)
	switch subChoice {
	case 1:
		fmt.Println("Linear Search")
		i := searchLinear(intSlice, target)
		fmt.Printf("found target %v at position %v in intSlice %v \n", target, i+1, intSlice)
	case 2:
		fmt.Println("Two pointer")
		i := searchTwoPointer(intSlice, target)
		fmt.Printf("found target %v at position %v in intSlice %v \n", target, i+1, intSlice)

	default:
		fmt.Println("not a valid sub choice")

	}

}

func searchLinear(intSlice []int, target int) int {
	for i := 0; i < len(intSlice); i += 1 {
		if intSlice[i] == target {
			return i
		}
	}
	return -1
}

func searchTwoPointer(intSlice []int, target int) int {
	for i, j := 0, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		if intSlice[i] == target {
			return i
		} else if intSlice[j] == target {
			return j
		}
	}
	return -1
}

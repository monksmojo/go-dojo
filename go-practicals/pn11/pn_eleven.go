package pn11

import (
	"fmt"
	"math"
)

/*
wap to find the second largest element in the slice of positive integers
for ex
input: int[]{11,12,40,50,23,24}
output:  40
*/

func FindSecondLargest() {
	numbers := []int{11, 12, 40, 50, 23, 24}
	fmt.Println("we use two ways to to find 2nd largest element in the slice")
	fmt.Println("1. using the default approach")
	fmt.Println("2. using the default + two pointers approach")
	fmt.Println("Choose the method to solve: ")
	var choice int
	fmt.Scanln(&choice)
	var r int
	switch choice {
	case 1:
		r = getSecondLargestV1(numbers)
	case 2:
		r = getSecondLargestV2(numbers)
	}
	fmt.Printf("the second largest integer in the slice %v is %v \n", numbers, r)
}

func getSecondLargestV1(intSlice []int) int {
	l, sl := intSlice[0], math.MinInt
	for i := 1; i < len(intSlice)-1; i++ {
		value := intSlice[i]
		if value > l {
			l, sl = intSlice[i], l
		} else if value < l && value >= sl {
			sl = value
		}
	}
	return sl
}

func getSecondLargestV2(intSlice []int) int {
	l, sl := intSlice[0], math.MinInt
	for i, j := 1, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		var value int
		if intSlice[i] > intSlice[j] {
			value = intSlice[i]
		} else {
			value = intSlice[j]
		}

		if value > l {
			l, sl = intSlice[i], l
		} else if value < l && value >= sl {
			sl = value
		}
	}
	return sl
}

package pn21

import (
	"fmt"
	"math"
)

func SecondSmallestInt() {
	var intArraySlice []int = []int{10, 7, 1, 2, 19, 0, 0, 1, 2, 10, 7}
	secondSmallestInt := fetchSecondSmallestInt(intArraySlice)
	fmt.Println("second smallest integer in the slice:: ", intArraySlice, " is ", secondSmallestInt)
	secondSmallestInt = fetchSecondSmallestIntUsingTwoPtr(intArraySlice)
	fmt.Println("second smallest integer in the slice:: ", intArraySlice, " is ", secondSmallestInt)

}

func fetchSecondSmallestInt(intArraySlice []int) int {
	var smallest int = math.MaxInt64
	var secondSmallest int = math.MaxInt64
	for i := 0; i < len(intArraySlice); i++ {
		if intArraySlice[i] < smallest {
			secondSmallest = smallest
			smallest = intArraySlice[i]
		} else if intArraySlice[i] < secondSmallest {
			secondSmallest = smallest
		}
	}
	return secondSmallest
}

func fetchSecondSmallestIntUsingTwoPtr(intArraySlice []int) int {
	var smallest int = math.MaxInt64
	var secondSmallest int = math.MaxInt64
	i, j := 0, len(intArraySlice)-1
	for i < j {
		if intArraySlice[i] < smallest {
			secondSmallest = smallest
			smallest = intArraySlice[i]
		} else if intArraySlice[i] < secondSmallest {
			secondSmallest = intArraySlice[i]
		}

		if intArraySlice[j] < smallest {
			secondSmallest = smallest
			smallest = intArraySlice[j]
		} else if intArraySlice[j] < secondSmallest {
			secondSmallest = intArraySlice[j]
		}

		i += 1
		j -= 1
	}
	return secondSmallest
}

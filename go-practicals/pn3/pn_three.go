package pn3

import (
	"fmt"
)

// Practical3 Description
/**
@initialAuthor: monks_mojo
@lastUpdatedBY: monks_mojo
Practical Three:
ASSIGNMENT
We've been asked to "bucket" costs for an entire month into the cost occurring on each day of the month.
Complete the getCostsByDay function. It should return a slice of float64s, where each element is the total cost for that day. The length of the slice should be equal to the number of days represented in the costs slice, including any days that have no costs, up to the last day represented in the slice.
Days are numbered and start at 0.

EXAMPLE
Input in day, cost format:

[] cost{
(0, 4.0),
{1, 2.1),
{1, 3.1),
{5, 2.5),
}

// here cost is a struct
type cost struct{
    day int
    value float64
}

// Output in day, slice of float64 format:
[]float64{4.0,5.2,0.0,0.0,0.0,2.5}


We use two approach to solve this problem
1. using make() function
2. using nested for loop

*/

// The function "CostByDay" allows the user to choose between two different methods for calculating the
// cost by day.
func CostByDay() {
	var choice int
	exercisePrompt := `We use two approach to solve this problem.
	1. using make() function.
	2. using nested for loop`
	fmt.Println(exercisePrompt)
	fmt.Println("Choose the approach")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("using make() function")
		costByDayType1(testCase1())
	case 2:
		fmt.Println("using nested for loop")
		costByDayType2(testCase1())
	}
}

// The type "cost" represents a cost value associated with a specific day.
// @property {int} day - An integer representing the day of the cost.
// @property {float64} value - The "value" property is a floating-point number that represents the cost
// value.
type cost struct {
	day   int
	value float64
}

// The function testCase1 returns a slice of cost structs with different day and value values.
func testCase1() []cost {
	return []cost{{day: 0, value: 4.0}, {day: 1, value: 2.1}, {day: 1, value: 3.1}, {day: 5, value: 2.5}}
}

// The function `costByDayType1` calculates the total cost per day based on a given slice of cost structures.
func costByDayType1(costStructSlice []cost) {
	maxDay := 0
	for _, costStruct := range costStructSlice {
		if costStruct.day > maxDay {
			maxDay = costStruct.day
		}
	}
	costPerDaySlice := make([]float64, maxDay+1)
	for _, costStruct := range costStructSlice {
		costPerDaySlice[costStruct.day] += costStruct.value
	}
	fmt.Println("cost per day is")
	fmt.Println(costPerDaySlice)
}

// The function calculates the cost per day based on a slice of cost structures.
func costByDayType2(costStructSlice []cost) {
	var costPerDaySlice = []float64{}
	for i := 0; i < len(costStructSlice); i += 1 {
		costStruct := costStructSlice[i]
		for costStruct.day >= len(costPerDaySlice) {
			costPerDaySlice = append(costPerDaySlice, 0.00)
		}
		costPerDaySlice[costStruct.day] += costStruct.value
	}
	fmt.Println("cost per day is")
	fmt.Println(costPerDaySlice)
}

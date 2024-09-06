package main

import "fmt"

func main() {
	fmt.Println("welcome to the switch sate")
	//  in Go, the switch statement terminates after the first matching case hence we don't use break keyword
	var daysOfWeek = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("enter current week number")
	var weekNumber int
	fmt.Scanln(&weekNumber)
	switch weekNumber {
	case 1:
		fmt.Println(daysOfWeek[0])
	case 2:
		fmt.Println(daysOfWeek[1])
	case 3:
		fmt.Println(daysOfWeek[2])
	case 4:
		fmt.Println(daysOfWeek[3])
	case 5:
		fmt.Println(daysOfWeek[4])
	case 6:
		fmt.Println(daysOfWeek[5])
	case 7:
		fmt.Println(daysOfWeek[6])
	default:
		fmt.Println("wrong week number inserted")
	}

	fmt.Println("go lang switch with multiple case")
	switch weekNumber {
	case 1, 2, 3, 4, 5:
		fmt.Println("oh no its a weekday")
	case 6, 7:
		fmt.Println("Yay! it's a weekend")
	}

	fmt.Println("go lang switch without an expression")
	switch {
	case 5 == weekNumber:
		fmt.Println("TGIF")
	default:
		fmt.Println("we still have time to reach to friday")
	}

}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to study the time package in golang")
	var presentTime time.Time = time.Now()
	fmt.Println("current time in default format", presentTime)
	// formatting current time in a format
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	currentDate := time.Date(2020, time.August, 12, 23, 59, 59, 0, time.UTC)
	fmt.Println(currentDate)
	fmt.Println("Formatted current date ", currentDate.Format("01-02-2006 15:04:05 Monday"))
}

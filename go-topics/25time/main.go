package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time package in Go")
	var defaultTime time.Time
	fmt.Println("the default value of time.Time: ", defaultTime)

	// printing current time
	var currentTime time.Time = time.Now()
	fmt.Println("current time is: ", currentTime)

	// creating a time using constructor
	// time.Date(year, month, day, hour, min, sec, nsec, loc).
	specificTime := time.Date(2023, 4, 1, 12, 0, 0, 0, time.Local)
	fmt.Println(specificTime)

	// extract hour from specific time
	specificHour := specificTime.Hour()
	fmt.Println("the current hour from the specific time is ", specificHour)

	// extract day of the month from specific time
	specificDay := specificTime.Day()
	fmt.Println("the current day from the specific time is ", specificDay)

	// extract day name of the month from specific time
	specificDayName := specificTime.Weekday().String()
	fmt.Println("the current day name from the specific time is ", specificDayName)

	// extract month from specific time
	// we use %d because by default prints month name
	specificMonth := specificTime.Month()
	fmt.Printf("the current month from the specific time is %d \n", specificMonth)

	// extract month name from specific time
	specificMonthName := specificTime.Month().String()
	fmt.Println("the current month name from the specific time is ", specificMonthName)

	// extract year from specific time
	specificYear := specificTime.Year()
	fmt.Println("the current year from the specific time is ", specificYear)

	specificTime = time.Date(2023, 4, 30, 23, 59, 59, 0, time.Local)
	fmt.Printf("new specific time is %v \n", specificTime)

	// adding 3 hour to specific time
	specificTime = specificTime.Add(3 * time.Hour)
	fmt.Printf("adding three hours to the specific time %v \n", specificTime)

	// subtracting 3 hour to specific time
	specificTime = specificTime.Add(-3 * time.Hour)
	fmt.Printf("subtracting three hours to the specific time %v \n", specificTime)

	// adding 3 days to specific time
	specificTime = specificTime.AddDate(0, 0, 3)
	fmt.Printf("adding three days to the specific time %v \n", specificTime)

	// subtracting 3 days to specific time
	specificTime = specificTime.AddDate(0, 0, -3)
	fmt.Printf("subtracting three days to the specific time %v \n", specificTime)

	// adding 3 months to specific time
	specificTime = specificTime.AddDate(0, 3, 0)
	fmt.Printf("adding three months to the specific time %v \n", specificTime)

	// subtracting 3 months to specific time
	specificTime = specificTime.AddDate(0, -3, 0)
	fmt.Printf("subtracting three months to the specific time %v \n", specificTime)

	// adding 3 years to specific time
	specificTime = specificTime.AddDate(3, 0, 0)
	fmt.Printf("adding three years to the specific time %v \n", specificTime)

	// subtracting 3 years to specific time
	specificTime = specificTime.AddDate(-3, 0, 0)
	fmt.Printf("subtracting three years to the specific time %v \n", specificTime)

	// formatting the date
	// The layout string uses a specific reference time: Mon Jan 2 15:04:05 MST 2006 to define the format.
	// "2006-01-02 15:04:05 MST" --> use this to format
	// You need to adjust your layout strings to match this reference time.
	layout1 := "02-01-2006"
	layout2 := "2006-01-02"
	layout3 := "02-Jan-2006"
	formattedCurrentDate := currentTime.Format(layout1)
	fmt.Println("formatting currentTime in format1", formattedCurrentDate)
	formattedCurrentDate = currentTime.Format(layout2)
	fmt.Println("formatting currentTime in format2", formattedCurrentDate)
	formattedCurrentDate = currentTime.Format(layout3)
	fmt.Println("formatting currentTime in format3", formattedCurrentDate)

	// converting string to date
	strDate := "2023-01-20"
	date, err := time.Parse(layout2, strDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fmt.Println(strDate, " string date converted to date type ", date.Format(layout2))

	fmt.Println("get last week time with hour 23:59:59")
	lastWeekDate, err := getLastWeekDate()
	if err != nil {
		errors.New("not abel to get last week date")
	} else {
		fmt.Println(lastWeekDate)
	}

	beforeTime := currentTime.Add(-5 * time.Hour)
	fmt.Println("5 hours before current time is ", beforeTime)
	diffInHours := getHours(currentTime, beforeTime)
	fmt.Println("difference in current time and before time is ", diffInHours)

}

func getLastWeekDate() (time.Time, error) {
	fmt.Println("objective is to get last week date from current date")
	dateLayout := "02-01-2006"
	dateTimeLayout := "02-01-2006 15:04:05"
	lastWeekDate := time.Now().AddDate(0, 0, -7).Format(dateLayout)
	lastWeekDate += "23:59:59"
	fmt.Println("last week formatted date", lastWeekDate)
	return time.Parse(dateTimeLayout, lastWeekDate)

}

func getHours(time1, time2 time.Time) float64 {
	return time1.Sub(time2).Hours()
}

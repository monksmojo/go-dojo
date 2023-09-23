package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// exception handling in go
	// error handling in the go
	// in go error act as the value
	fmt.Println("the errors in go")

	var err error
	fmt.Println(err) // default value of error is nil

	value, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println(" cannot covert string to int", err.Error())
	} else {
		fmt.Printf("%v %T", value, value)
	}

	/**
	when something can go wrong in a function
	the function must return error as the last value
	**/
	testMsg()

	/**
	calling a function that implements the custom error
	**/
	u1 := user{
		name:     "Gill",
		age:      -90,
		errorMsg: "",
	}
	printResult(u1)

	u2 := user{
		name:     "",
		age:      90,
		errorMsg: "",
	}
	printResult(u2)

	u3 := user{
		name:     "Mai",
		age:      70,
		errorMsg: "",
	}
	printResult(u3)

	// type assertion with error interface
	typeAssertionWithError()

}

/*
* write a go program that that send message to couple
and calculate cost of the message
the function will throw error if the length of the msg >25
*
*/

func testMsg() {
	totalCost, err := sendMsgToCouple("your gift has been delivered", "happy birthday")
	if err != nil {
		fmt.Printf("%v hence cost is %.2f \n", err, totalCost)
	} else {
		fmt.Printf("total cost for sending the message are %.2f \n", totalCost)
	}
}

func sendMsgToCouple(msg, spouseMsg string) (float64, error) {
	var msgCost, spouseMsgCost float64
	var err error
	msgCost, err1 := calculateSmSCost(getMsgLength(msg))
	spouseMsgCost, err2 := calculateSmSCost(getMsgLength(spouseMsg))
	if err1 != nil && err2 != nil {
		err = fmt.Errorf("%v while sending to the customer and its spouse", err1)
	} else if err1 != nil {
		err = fmt.Errorf("%v while sending to the customer", err1)
	} else if err2 != nil {
		err = fmt.Errorf("%v while sending to the spouse", err2)
	}
	return msgCost + spouseMsgCost, err

}

func calculateSmSCost(msgLength float64) (cost float64, err error) {
	if msgLength > 25.0 {
		cost = 0.00
		err = fmt.Errorf("the length of the message exceeds 25 characters")
	} else {
		cost = 0.02 * msgLength
	}
	return cost, err
}

func getMsgLength(msg string) float64 {
	return float64(len(msg))
}

/*
*
Creating custom error in go
*
*/
func printResult(u user) {
	ft, customErr := testCustomError(u)
	if customErr != nil {
		fmt.Println(customErr.Error())
	} else {
		fmt.Println(ft.GetMsg())
	}
}

func testCustomError(u user) (format, error) {
	return validateUser(u)
}

func validateUser(u user) (format, error) {
	if u.name == "" {
		u.errorMsg = "name can't be blank"
		return nil, u
	} else if u.age < 0 {
		u.errorMsg = "age can't be negative"
		return nil, u
	} else {
		return u, nil
	}
}

// // an error is type that implements the error interface
type error interface {
	Error() string
}

type format interface {
	GetMsg() string
}

type user struct {
	name     string
	age      int
	errorMsg string
}

func (u user) Error() string {
	return u.errorMsg
}

func (u user) GetMsg() string {
	return fmt.Sprintf("welcome a %d year old %v to the club \n", u.age, u.name)
}

// The function "divide" takes two float64 numbers as input and returns their division as a float64,
// along with an error if the denominator is zero.
func divide(numerator, denominator float64) (result float64, err error) {
	if denominator == 0.00 {
		return 0.00, errors.New("cant divide by 0")
	}
	return numerator / denominator, nil
}

// The function `typeAssertionWithError` demonstrates the use of type assertions and error handling in Go.
func typeAssertionWithError() {
	u1 := user{
		name:     "Gill",
		age:      -90,
		errorMsg: "",
	}
	_, err1 := testCustomError(u1)
	logger(err1)
	_, err2 := divide(45.0, 0.00)
	logger(err2)
}

// The logger function prints out the error message based on the type of error received.
func logger(err error) {
	switch v := err.(type) {
	case user:
		fmt.Println("error with the userStruct")
		fmt.Println(v.Error())
	default:
		fmt.Println(v.Error())
	}
}

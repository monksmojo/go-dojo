package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to GOLANG !"
	fmt.Println("Hey ", welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings of GOLANG ! language so far !")

	// comma ok syntax is use in go as a alternative to try and catch
	// there is not really any try catch in go golang.
	// the exception is return by the function  in the form of variable
	// the exception in golang is received like an variable
	userRating, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating", userRating)
	fmt.Printf("type of the rating is %T", userRating)
	// the _ in the line 19 can be replaced by the variable name
	// if we want to handle the error
	// but if we don't want to handle the error we can ignore the error by _

	// different ways of taking input from user in GO

	// 1.bufio.NewReader(os.Stdin)
	// taking input using bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to our pizza parlour")
	fmt.Println("please rate our pizza from 1 to 5")
	ratingReader := bufio.NewReader(os.Stdin)
	userInput, _ := ratingReader.ReadString('\n')
	fmt.Println("Thanks for the rating: ", userInput)
	fmt.Printf("rating : is of Type %T \n", userInput)

	// 2.bufio.NewScanner(os.Stdin)
	// taking input using bufio.NewScanner(os.Stdin)

	fmt.Println("welcome to bookCab")
	fmt.Println("please rate our application from 1 to 5")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput = scanner.Text()
	fmt.Println("Thanks for the rating: ", userInput)
	fmt.Printf("rating : is of Type %T \n", userInput)

	// 3. fmt.Scan()
	// the Scan() function takes input from the user. However, it can only take input values up to a space.
	// When it encounters a space, the value before space is assigned to the specified variable. For example,
	fmt.Println("welcome to han solo arcade")
	fmt.Println("Rate our arcade! ")
	var rating float64
	fmt.Scan(&rating)
	fmt.Println("Thanks !! rating of the arcade is ", rating)
	fmt.Printf("rating data type is %T \n", rating)

	// taking multiple input using fmt.Scan()
	var name string
	var age int
	fmt.Println("please enter you name and ag>e")
	_, error := fmt.Scan(&name, &age)
	if error != nil {
		fmt.Println("an error occurred while taking your name and age")
		fmt.Println(error)
	} else {
		fmt.Printf("your name is %s and you age is %d\n", name, age)
		fmt.Printf("your name is of data type %T and you age is of data type %T\n", name, age)
	}

	//4. fmt.Scanln()
	// We use the Scanln() function to get input values up to the new line. When it encounters a new line, it stops taking input values. For example,
	var gameName string
	var duration int
	fmt.Println("enter the game you want to play !")
	fmt.Scanln(&gameName)
	fmt.Println("enter the duration in minutes !")
	fmt.Scanln(&duration)
	fmt.Printf("game you want to play is %s and the duration is %d mins \n", gameName, duration)
	fmt.Printf("gameName data type is %T and the duration is of data type %T mins \n", gameName, duration)

	//5. fmt.scanf()
	// the fmt.Scanf() function is similar to fmt.Scanln() , the only difference is it takes input using  format specifier()
	var score float64
	var isHighScore bool
	fmt.Println("enter your score")
	fmt.Println("is it highScore")
	fmt.Scanf("%f %t", &score, &isHighScore)
	fmt.Println("your score is ", score, " is it high score", isHighScore)

	// fmt.Scan() fmt.Scanln() returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.
}

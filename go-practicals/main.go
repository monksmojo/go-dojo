package main

import (
	"fmt"

	"github.com/monksmojo/go-dojo/go-practicals/pn1"
	"github.com/monksmojo/go-dojo/go-practicals/pn2"
	"github.com/monksmojo/go-dojo/go-practicals/pn3"
	"github.com/monksmojo/go-dojo/go-practicals/pn4"
	"github.com/monksmojo/go-dojo/go-practicals/pn5"
	"github.com/monksmojo/go-dojo/go-practicals/pn6"
)

func main() {
	fmt.Println("Welcome To The GoLang Practical File")
	fmt.Println("A menu based driven practical file")
	choice := "y"
	for choice == "y" || choice == "Y" {
		fmt.Println("List of practicals in the file")
		fmt.Println("PN :1 print all the prime numbers up to the given nTh integer")
		fmt.Println("PN :2 reverse the element of the array using")
		fmt.Println("PN :3 cost by day from video Go Programming : Golang Course with Bonus Projects 3:23:00")
		fmt.Println("PN :4 2D matrix printing from video Go Programming : Golang Course with Bonus Projects 3:29:30")
		fmt.Println("PN :5 Creating a name dictionary which holds names with its frequency Go Programming : Golang Course with Bonus Projects 4:05:22")
		fmt.Println("PN :6 Take a slice and filter all the multiples of 3 and 5 from the slice")
		fmt.Println("Select the practical you want to execute")
		var pn int
		fmt.Scanln(&pn)
		switch pn {
		case 1:
			fmt.Println("PN :1 print all the prime numbers up to the given nTh integer")
			pn1.CalculatePrime()
		case 2:
			fmt.Println("PN :2 reverse the element of the array using")
			pn2.ReverseArray()
		case 3:
			fmt.Println(" :3 cost by day from video Go Programming : Golang Course with Bonus Projects 3:23PN")
			pn3.CostByDay()
		case 4:
			fmt.Println("PN :4 2D matrix printing from video Go Programming : Golang Course with Bonus Projects 3:29:30")
			pn4.CreateMatrix()
		case 5:
			fmt.Println("PN :5 Creating a name dictionary which holds names with its frequency Go Programming : Golang Course with Bonus Projects 4:05:22")
			pn5.PrintNameCounts()
		case 6:
			fmt.Println("PN :6 Take a slice and filter all the multiples of 3 and 5 from the slice")
			pn6.FilterSlice()
		default:
			fmt.Println("you entered the wrong choice dummy !")
		}
		fmt.Println("Press 'y' to execute more practicals press 'n' to exit.")
		fmt.Scanln(&choice)
	}
	fmt.Println("BYE BYE !")
}

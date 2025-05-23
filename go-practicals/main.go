package main

import (
	"fmt"

	"github.com/monksmojo/go-dojo/go-practicals/pn1"
	"github.com/monksmojo/go-dojo/go-practicals/pn10"
	"github.com/monksmojo/go-dojo/go-practicals/pn11"
	"github.com/monksmojo/go-dojo/go-practicals/pn12"
	"github.com/monksmojo/go-dojo/go-practicals/pn13"
	"github.com/monksmojo/go-dojo/go-practicals/pn14"
	"github.com/monksmojo/go-dojo/go-practicals/pn15"
	"github.com/monksmojo/go-dojo/go-practicals/pn16"
	"github.com/monksmojo/go-dojo/go-practicals/pn17"
	"github.com/monksmojo/go-dojo/go-practicals/pn18"
	"github.com/monksmojo/go-dojo/go-practicals/pn19"
	"github.com/monksmojo/go-dojo/go-practicals/pn2"
	"github.com/monksmojo/go-dojo/go-practicals/pn20"
	"github.com/monksmojo/go-dojo/go-practicals/pn21"
	"github.com/monksmojo/go-dojo/go-practicals/pn22"
	"github.com/monksmojo/go-dojo/go-practicals/pn3"
	"github.com/monksmojo/go-dojo/go-practicals/pn4"
	"github.com/monksmojo/go-dojo/go-practicals/pn5"
	"github.com/monksmojo/go-dojo/go-practicals/pn6"
	"github.com/monksmojo/go-dojo/go-practicals/pn7"
	"github.com/monksmojo/go-dojo/go-practicals/pn8"
	"github.com/monksmojo/go-dojo/go-practicals/pn9"
)

func main() {

	fmt.Println("Welcome To The GoLang Practical File")
	fmt.Println("A menu based driven practical file")
	choice := "y"
	for choice == "y" || choice == "Y" {
		fmt.Println("List of practicals in the file")
		fmt.Println("PN:1 print all the prime numbers up to the given nTh integer")
		fmt.Println("PN:2 reverse the element of the array using")
		fmt.Println("PN:3 cost by day from video Go Programming : Golang Course with Bonus Projects 3:23:00")
		fmt.Println("PN:4 2D matrix printing from video Go Programming : Golang Course with Bonus Projects 3:29:30")
		fmt.Println("PN:5 Creating a name dictionary which holds names with its frequency Go Programming : Golang Course with Bonus Projects 4:05:22")
		fmt.Println("PN:6 Take a slice and filter all the multiples of 3 and 5 from the slice")
		fmt.Println("PN:7 Write a Go function to swap the consecutive characters in a string input by the user.")
		fmt.Println("PN:8 Write a Go function to input 2 strings i.e. s1 and s2 and find the presence of s1 in s2. and encode it")
		fmt.Println("PN9: Write a Go function to remove the characters which have odd index values in a given string.")
		fmt.Println("PN10: Write a Go function to find frequency of each character 	of the string.")
		fmt.Println("PN11: Write a Go function to find the second largest element in the slice")
		fmt.Println("PN:12 Write a Go function to convert a given string to all uppercase if it contains at least uppercase characters in the first 4 characters.")
		fmt.Println("PN:13 Write a Go function to remove duplicate characters of a given string.")
		fmt.Println("PN:14 Write a Go function to print the reverse of the string and check for palindrome")
		fmt.Println("PN:15 Write a Go function to get a single string from two given strings, separated by a space and swap the first two characters of each string. Eg input: 'abc', 'xyz' output: 'xyc abz'")
		fmt.Println("PN:16 LeetCode 744. Find Smallest Letter Greater Than Target")
		fmt.Println("PN:17 GeekForGeek: Find uncommon characters of the two strings")
		fmt.Println("PN:18 LinearSearch: search an int in slice of int")
		fmt.Println("PN:19 LinearSearch: search a character in slice of characters")
		fmt.Println("PN:20 LinearSearch: search a word in a string")
		fmt.Println("PN:21 LinearSearch: search for second smallest int in int slice")
		fmt.Println("PN:22 LinearSearch: search for  int in 2D matrix")
		fmt.Println("PN:23 LinearSearch: search for second largest element in 2D matrix")
		fmt.Println("PN:24 LinearSearch: extract numbers with even digits from a slice of int ")
		fmt.Println("Select the practical you want to execute")
		var pn int
		fmt.Scanln(&pn)
		switch pn {
		case 1:
			fmt.Println("PN:1 print all the prime numbers up to the given nTh integer")
			pn1.CalculatePrime()
		case 2:
			fmt.Println("PN:2 reverse the element of the array")
			pn2.ReverseArray()
		case 3:
			fmt.Println("PN:3 cost by day from video Go Programming : Golang Course with Bonus Projects 3:23PN")
			pn3.CostByDay()
		case 4:
			fmt.Println("PN:4 2D matrix printing from video Go Programming : Golang Course with Bonus Projects 3:29:30")
			pn4.CreateMatrix()
		case 5:
			fmt.Println("PN:5 Creating a name dictionary which holds names with its frequency Go Programming : Golang Course with Bonus Projects 4:05:22.")
			pn5.PrintNameCounts()
		case 6:
			fmt.Println("PN:6 Take a slice and filter all the multiples of 3 and 5 from the slice.")
			pn6.FilterSlice()
		case 7:
			fmt.Println("PN:7 Write a Go function to swap the consecutive characters in a string input by the user.")
			pn7.SwapCharactersOfString()
		case 8:
			fmt.Println("PN:8 Write a Go function to input 2 strings i.e. s1 and s2 and find the presence of s1 in s2 and encode it.")
			pn8.EncodeTheString()
		case 9:
			fmt.Println("PN:9 Write a Go function to remove the characters which have odd index values in a given string.")
			pn9.FilterOddIndex()
		case 10:
			fmt.Println("PN:10 Write a Go function to find frequency of each character 	of the string.")
			pn10.CalculateFunction()
		case 11:
			fmt.Println("PN:11 Write a Go function to find the second largest element in the slice")
			pn11.FindSecondLargest()
		case 12:
			fmt.Println("PN:12 Write a Go function to convert a given string to all uppercase if it contains at least uppercase characters in the first 4 characters.")
			pn12.StrToUpper()
		case 13:
			fmt.Println("PN:13 Write a Go function to remove duplicate characters of a given string.")
			pn13.RemoveDuplicates()
		case 14:
			fmt.Println("PN:14 Write a Go function to print the reverse of the string and check for palindrome")
			pn14.ReverseTheString()
		case 15:
			fmt.Println("PN:15 Write a Go function to get a single string from two given strings, separated by a space and swap the first two characters of each string. Eg input: 'abc', 'xyz' output: 'xyc abz'")
			pn15.CreateString()
		case 16:
			fmt.Println("PN:16 LeetCode 744. Find Smallest Letter Greater Than Target")
			pn16.GetNextGreatestLetter()
		case 17:
			fmt.Println("PN:17 GeekForGeek: Find uncommon characters of the two strings")
			pn17.FindUncommonCharacter()
		case 18:
			fmt.Println("PN:18 LinearSearch: search an int in slice of int")
			pn18.SearchInIntSlice()
		case 19:
			fmt.Println("PN:19 LinearSearch: search a character in slice of characters")
			pn19.SearchCharacterSlice()
		case 20:
			fmt.Println("PN:20 LinearSearch: search a word in a string")
			pn20.SearchWordInString()
		case 21:
			fmt.Println("PN:21 LinearSearch: search for second smallest int in int slice")
			pn21.SecondSmallestInt()
		case 22:
			fmt.Println("PN:22 LinearSearch: search for  int in 2D matrix")
			pn22.SearchIn2DMatrix()
		case 23:
			fmt.Println("PN:23 LinearSearch: search for second largest element in 2D matrix")
		case 24:
			fmt.Println("PN:24 LinearSearch: extract numbers with even digits from a slice of int ")
		case 25:
			fmt.Println("PN:25 DeepEdge AI:  a loyal customer")
		default:
			fmt.Println("you entered the wrong choice dummy !")
		}
		fmt.Println("Press 'y' to execute more practicals press 'n' to exit.")
		fmt.Scanln(&choice)
	}
	fmt.Println("BYE BYE !")
}

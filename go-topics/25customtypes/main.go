package main

import "fmt"

/*
    Basic Type: int , float64 , byte , rune , string , boolean , complex , etc.

	Composite Type (Non-Reference): Arrays , Structs

	Composite Type (Reference): Slices , Maps , Channel , Pointer , Function/Method

	Interface: interface
*/
/*
But besides all of the built-in data types, we can create our data type by simply using type .
*/
type age int

type person struct {
	personName string
	personAge  age
}

func main() {

	/*
		custom types are types aliases or new type definitions created from existing types
		which can be functions too
	*/
	p := person{
		personName: "Peter",
		personAge:  age(22),
	}

	fmt.Printf("is %v eligible to vote %v", p.personName, eligibleToVote(p.personAge))

}

func eligibleToVote(personAge age) bool {
	return int(personAge) >= 18
}

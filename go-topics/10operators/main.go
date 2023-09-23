package main

import "fmt"

func main() {
	fmt.Println("Welcome to Operators in GO!")
	// Arithmetic Operator
	// We use arithmetic operators to perform arithmetic operations like addition, subtraction, multiplication, and division.
	// + (Addition)	a + b
	// - (Subtraction)	a - b
	// *  (Multiplication)	a * b
	// / (Division)	a / b
	// % (Modulo Division)	a % b
	a := 10
	b := 35
	fmt.Printf("addition of %d and %d is %d \n", a, b, a+b)
	fmt.Printf("subtraction of %d and %d is %d \n", b, a, b-a)
	fmt.Printf("multiplication of %d and %d is %d \n", b, a, b*a)
	fmt.Printf(" division of %d and %d is %d \n", b, a, b/a)

	// to add an integer and float64 in go we have to use type casting.
	//  golang does not support adding of int + float64 without type casting
	var c int = 90
	var d float64 = 90.05
	e := float64(c) + d
	fmt.Printf("addition of %d+%f = %f", c, d, e)

	// same goes with subtraction, multiplication division and modulo

	// when we divide a int with int we get a int as a result as seen in statement above 35/10 will give 3

	f := 10.0
	g := 15.0
	fmt.Printf(" division of %f and %f is %f \n", g, f, g/f)

	// increment operator ++
	// decrement operator --
	/*
		only post ++ and post -- operator are supported in go
	*/
	var x int = 10
	x++
	fmt.Println("post increment", x)
	x--
	fmt.Println("post decrement", x)

	// assignment operator and combined assignment operator
	x = 90
	fmt.Println("x=90 value of x is assigned as ", x)
	// += (addition assignment)	a += b means	a = a + b
	x += 5
	fmt.Println(x)
	// -= (subtraction assignment)	a -= b means	a = a - b
	x -= 5
	fmt.Println(x)
	// *= (multiplication assignment)	a *= b means	a = a * b
	x *= 5
	fmt.Println(x)
	// /= (division assignment)	a /= b means	a = a / b
	x /= 5
	fmt.Println(x)
	// %= (modulo assignment)	a %= b means a = a % b
	x %= 5
	fmt.Println(x)

	// comparison and relational operator in go
	/*
		== (equal to)	a == b	returns true if a and b are equal
		!= (not equal to)	a != b	returns true if a and b are not equal
		> (greater than)	a > b	returns true if a is greater than b
		< (less than)	a < b	returns true if a is less than b
		>= (greater than or equal to)	a >= b	returns true if a is either greater than or equal to b
		<= (less than or equal to)	a <= b	returns true is a is either less than or equal to b
	*/
	var w float64 = 90.0
	var y float64 = 90.76
	var z float64 = 100
	fmt.Printf("%f == %f is %t \n", w, y, w == y)
	fmt.Printf("%f != %f is %t \n", w, y, w != y)
	fmt.Printf("%f > %f is %t \n", y, w, y > w)
	fmt.Printf("%f >= %f is %t \n", y, w, y >= w)
	fmt.Printf("%f < %f is %t \n", y, w, y < w)
	fmt.Printf("%f <= %f is %t \n", y, w, y <= w)

	// logical operator in go language
	/*
		&& (Logical AND)	exp1 && exp2	returns true if both expressions exp1 and exp2 are true

		|| (Logical OR)	exp1 || exp2	returns true if any one of the expressions is true.

		! (Logical NOT)	!exp	returns true if exp is false and returns false if exp is true.
	*/

	var entry bool = !(w == z && z < y)
	fmt.Printf("!(%f == %f && %f < %f is %t) \n", w, z, z, y, entry)

}

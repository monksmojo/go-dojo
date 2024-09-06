package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("welcome to the interface")
	/**
	interface are collection of method signatures
	**/
	/**
	a struct type implements the interface if it has all the
	method signature of the interface implemented in it
	**/

	rec := rectangle{
		shapes: shapes{
			shapeName: "Rectangle",
		},
		length:  40.25,
		breadth: 20.25,
	}

	var areaOfRectangle float64 = calculateArea(rec)
	fmt.Printf("%v area is %f \n", rec.shapeName, areaOfRectangle)

	var perimeterOfRectangle float64 = calculatePerimeter(rec)
	fmt.Printf("%v perimeter is %f \n", rec.shapeName, perimeterOfRectangle)

	cer := circle{
		shapes: shapes{
			shapeName: "Circle",
		},
		radius: 20.31,
	}

	var areaOfCircle float64 = calculateArea(cer)
	fmt.Printf("%v area is %f \n", cer.shapeName, areaOfCircle)

	var perimeterOfCircle float64 = calculatePerimeter(cer)
	fmt.Printf("%v perimeter is %f \n", cer.shapeName, perimeterOfCircle)
}

/*
*
creating an interface shape with area and parameter as
method signatures.
*
*/
type shape interface {
	area() float64
	perimeter() float64
}

/*
*
creating a struct rectangle that implements this interface
*
*/
type shapes struct {
	shapeName string
}

type rectangle struct {
	shapes
	length  float64
	breadth float64
}

// implements the method signature area of the
func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length * r.breadth)
}

type circle struct {
	shapes
	radius float64
}

// The `func (c circle) area() float64` is a method implementation
// 'for the `area()` method of the `shape` interface for the `circle` struct.
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// The `func (c circle) perimeter() float64` is a method implementation for the `perimeter()` method of
// the `shape` interface for the `circle` struct.
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// The calculateArea function takes a shape interface as input argument and returns its area.
// so we can pass any type struct implementing shape interface method as function argument
func calculateArea(s shape) float64 {
	return s.area()
}

// The function "calculatePerimeter" shape interface as input argument and returns perimeter.
// so we can pass any type struct implementing shape interface method as function argument
func calculatePerimeter(s shape) float64 {
	return s.perimeter()
}

package customPackage

import "math"

var MathPackageName string = "this is a custom math package"



/**
some custom methods of math package
Functions	Descriptions
Sqrt()	returns the square root of the number
Cbrt()	returns the cube root of the number
Max()	returns the larger number between two
Min()	returns the smaller number between two
Mod()	computes the remainder after division
**/

var num1 float64 = 25.0

var num2 float64 = 125.0

func GetNum1() float64 {
	return num1
}

func GetNum2() float64 {
	return num2
}

func GetSquareRoot(num float64) float64 {
	return math.Sqrt(num)
}

func GetCubeRoot(num float64) float64 {
	return math.Cbrt(num)
}

func GetMaxOf2(num1 float64, num2 float64) float64 {
	return math.Max(num1, num2)
}

func GetMinOf2(num1 float64, num2 float64) float64 {
	return math.Min(num1, num2)
}

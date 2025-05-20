package main

import (
	"03package/customPackage"
	"fmt"
	"strings"
)

// the go language is a modular language where the code can be distributed  into different packager and can be compiled together into single binary
// one package can be imported into another using `import`` keyword

// importing single package
// import "fmt"

// importing multiple packages
/*
import (
	"fmt"
	"os"
	"strings"
)
*/

func main() {
	fmt.Println("welcome to main driver class")
	fmt.Println("calling customMath.go function", customPackage.MathPackageName)
	var num1 float64 = customPackage.GetNum1()
	var num2 float64 = customPackage.GetNum2()
	fmt.Println("float number we got from the number 1 is ", num1)
	fmt.Println("float number we got from the number 2 is ", num2)
	fmt.Printf("square root of %v is %v \n", num1, customPackage.GetSquareRoot(num1))
	fmt.Printf("cube root of %v is %v \n", num2, customPackage.GetCubeRoot(num2))
	fmt.Printf("greater of two number %v and %v is %v \n", num1, num2, customPackage.GetMaxOf2(num1, num2))
	fmt.Printf("minimum of two number %v and %v is %v \n", num1, num2, customPackage.GetMinOf2(num1, num2))
	fmt.Println()
	fmt.Println("calling string.go function", customPackage.StrPackageName)
	var string1 string = customPackage.GetStr1()
	var string2 string = customPackage.GetStr2()
	var string3 string = customPackage.GetStr3()
	slice1 := customPackage.GetSlice()
	fmt.Println("values fetched from package are", string1, string2, string3, slice1)
	fmt.Println(customPackage.Compare2Str(string1, string2))
	fmt.Printf("'red' occurred %v times in %v \n", customPackage.IfContains("red", string3), string3)
	fmt.Println(customPackage.CountTheOccurrence(string3, "red"))
	fmt.Printf("%v converted to %v \n", slice1, customPackage.ArrayToString(slice1))
	fmt.Printf("%v converted to upper case %v \n", string1, customPackage.StringInCase(string1, "upper"))
	fmt.Printf("%v converted to lower case %v \n", strings.ToUpper(string2), customPackage.StringInCase(string2, "lower"))

}

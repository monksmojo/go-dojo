package main

import (
	"fmt"
	"math"
)

/**
in GO
data type variables with public visibility across various packages
start with an uppercase
and data types and variables with private visibility to be used inside same package
starts with lowercase
**/

func main() {
	fmt.Println("Welcome to the structs in go programming")
	/**
	the structs in go are the user defined data type
	structs in go contain multiple data types called properties
	we can also bind function/methods to the respective structs

	a struct is defined by
	type keyword followed by struct name and ending with struct keyword
	type <struct_name> struct{
		DataTypeName datatype
		DataTypeName1 datatype1
	}

	they are defined outside main function
	**/

	// how to create a instance of the struct
	// method-1 --> dictionary method
	var m1 messageToSend = messageToSend{
		phoneNumber: 8877449955,
		message:     "Hey I am a struct m1",
	}
	fmt.Println(m1)

	//method-2 --> constructor method
	var m2 messageToSend = messageToSend{8844551122, "Hey I am a struct m2"}
	fmt.Println(m2)

	// method-3 --> declare and initialize with . operator
	var m3 messageToSend
	/** when we declare struct
		all its primitive data type
		properties are declared
		and initialized with default values
	**/
	m3.phoneNumber = 4455221166
	m3.message = "Hey I am a struct m3"
	fmt.Println(m3)

	// method 4 --> declare with {} brackets
	// and initialize with . operator
	m4 := messageToSend{}
	m4.phoneNumber = 5544774499
	m4.message = "Hey I am a struct m4"
	fmt.Println(m4)

	// method 5 --> declaring and initializing nested struct
	receipt1 := receipt{}
	receipt1.receiptNumber = 121
	receipt1.cost = 4500.54
	receipt1.shopInfo.shopName = "Get It On"

	// declaring and initializing person struct
	person1 := person{name: "John", age: 23}
	fmt.Println("a struct of type person", person1)

	// declaring a initializing person struct using . operator
	person2 := person{}
	person2.name = "Paul"
	person2.age = 24
	fmt.Println("a struct of type person", person2)

	// anonymous structs
	/**
	anonymous struct
	the struct which are declared and initialized once
	they don't have a name
	therefore cannot be referenced elsewhere in the code
	and are declared inside the function
	they are directly assigned to the variable
	**/

	cred := struct {
		url      string
		userName string
		password string
	}{
		url:      "http://Localhost:8080/dev/go",
		userName: "duke",
		password: "duke#qwerty@123",
	}
	fmt.Println("printing  an anonymous struct keywords", cred)

	// nested struct is when we declare a struct inside another struct
	// they are also an individual struct and can act independently
	var nvidiaGPU Processor = Processor{core_number: 4, processor_type: "GPU"}
	// declaring and initializing a nested struct
	pc1 := pc{
		brand:      "Lenovo",
		model:      "E590",
		processor1: nvidiaGPU,
		Processor2: Processor{
			core_number:    4,
			processor_type: "CPU",
		},
	}
	fmt.Println(" a pc struct with nested processor struct", pc1)

	/**
	embedded structs
	embedded structs help you to achieve oops composition
	here a parent struct is embedded into the child struct
	the fields of the embedded(parent) struct becomes the part of the child struct
	you can access the fields directly
	**/
	pwcProject := Project{
		projectName:       "PWC_WEB",
		projectDepartment: "Web Development",
	}

	employee := employee{
		empId: 2341,
		person: person{
			name: "Sid Sha",
			age:  21,
		},
		project1: Project{
			projectName:       "InZone",
			projectDepartment: "Mobile Apps",
		},
		project2: pwcProject,
	}

	fmt.Println(employee)

	// declaring a square struct and calling its methods
	s1 := square{
		shape: shape{
			shapeName:     "square",
			hasSides:      true,
			numberOfSides: 4,
		},
		length: 16.05,
	}

	var shapeName string = s1.GetShapeName()
	fmt.Printf("the shape name is %v, \n", shapeName)
	if s1.hasSides {
		sidesNumber := s1.numberOfSides
		fmt.Printf(" number of sides %v has are %d \n", shapeName, sidesNumber)
	}
	fmt.Printf("the area of %v is %f \n and perimeter of %v is %f", shapeName, s1.GetArea(), shapeName, s1.GetPerimeter())

	// declaring s struct with function inside as a field
	s2 := circle{
		shape: shape{
			shapeName:     "circle",
			hasSides:      false,
			numberOfSides: 0,
		},
		radius: 20.14,
		area: func(radius float64) (area float64) {
			return math.Pi * (radius * radius)
		},
	}
	fmt.Printf("the radius of %v is %f \n", s2.shapeName, s2.radius)
	fmt.Printf("the area of %v is %f \n", s2.shapeName, s2.area(s2.radius))

	// using pointers with the struct
	// to create struct properties getter and setter
	var angela person
	angela.SetName("Angela")
	angela.SetAge(27)
	fmt.Println(angela)

}

type messageToSend struct {
	phoneNumber int
	message     string
}

/*
*
nested struct we can declare a struct inside another struct
receipt is a struct which has Shop as a nested struct
*
*/
type receipt struct {
	receiptNumber int
	cost          float64
	shopInfo      shop
}

type shop struct {
	shopName string
}

type person struct {
	name string
	age  int
}

// using pointers with the struct
// to create struct properties getter and setter

func (p *person) SetName(name string) {
	p.name = name
}

func (p *person) SetAge(age int) {
	p.age = age
}

type employee struct {
	person
	empId    int
	project1 Project
	project2 Project
}

// project struct act as a nested struct to the  Employee Struct
type Project struct {
	projectName       string
	projectDepartment string
}

type Processor struct {
	core_number    int
	processor_type string
}

/*
*
pc is a struct with
with Processor nested struct
*
*/
type pc struct {
	brand      string
	model      string
	processor1 Processor
	Processor2 Processor
}

type shape struct {
	shapeName     string
	hasSides      bool
	numberOfSides int64
}

type square struct {
	shape
	length float64
}

func (s square) GetArea() (area float64) {
	area = s.length * s.length
	return area
}

func (s square) GetPerimeter() float64 {
	return float64(s.numberOfSides) * s.length
}

func (s square) GetShapeName() string {
	return s.shapeName
}

func (s square) GetSidesNumber() int64 {
	if s.hasSides {
		return s.numberOfSides
	}
	return 0
}

/*
*
Go also allow us to declare a function signature without defining its body
which later can defined to the function definition
*
*/
type CircleArea func(radius float64) (area float64)

/**
Go allow a function to be declared inside a struct as a field of the struct
**/

type circle struct {
	shape
	radius float64
	area   CircleArea
}

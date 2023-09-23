package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to maps in GO")
	// maps data structure in go store values in the form of key value pair
	// keys of the map are unique identifiers
	// where as values associated with the key can be duplicate

	// creation of map in go lang method-1
	// method 1 subject and marks
	// declaration and initialization
	subjectMarks := map[string]float32{"reasoning": 80.0, "python": 90.5, "go": 95.5}
	fmt.Println("subject and marks map ", subjectMarks)

	// creation of map using make
	// also adding key and value to the map in GO
	subjectGrade := make(map[string]string)
	subjectGrade["maths"] = "A"
	subjectGrade["python"] = "A+"
	subjectGrade["go"] = "O"
	fmt.Println(subjectGrade)

	// accessing values of the map in golang
	// we will be accessing the values
	colorMap := map[int]string{1: "red", 2: "blue", 3: "green"}
	color2, ok := colorMap[2]
	fmt.Printf("is the value present %t \n", ok)
	var msg string = fmt.Sprintf("the color we found on the second key is %s \n", color2)
	fmt.Println(msg)

	/*
		when accessing the values from the map the map returns
		two values one is the value bind to the key
		and another is boolean value [true/false]
		telling if key exist in the map
	*/

	// adding the value to the map
	colorMap[4] = "yellow"
	colorMap[5] = "peach"
	colorMap[6] = "violet"
	fmt.Println("new key value pair in map", colorMap)

	// updating the key's value in map
	colorMap[5] = "violet"
	fmt.Println("updated color map", colorMap)

	// deleting the key value pair from the map
	delete(colorMap, 6)
	fmt.Println("color map after deleting key value pair", colorMap)

	// in go maps are passed as reference to the function
	// a pointer pointing to the memory location
	ageMap := make(map[string]int)
	fmt.Printf("the value of map is %v memoryLocation %p before the function call \n", ageMap, ageMap)
	addValueToMap(ageMap)
	fmt.Printf("the value of map is %v memoryLocation %p after the function call \n", ageMap, &ageMap)

	/*
		only the comparable type can become the key of the maps in go
		boolean
		numeric
		string
		pointer
		channel
		interface
		all the above type can become key of the maps
		array,slices,map function type cannot become the key of the maps
	*/

	// creating a nested map
	var hits= make(map[string]map[string]int)
	addValueToMap1(hits, "google.com", "India")
	addValueToMap1(hits, "google.com", "Australia")
	addValueToMap1(hits, "Facebook.com", "India")
	addValueToMap1(hits, "Instagram.com", "India")
	fmt.Printf("the value of hitsMap is %v memoryLocation %p after the function call \n", hits, &hits)

}

func addValueToMap(aMap map[string]int) {
	aMap["yuri"] = 26
	fmt.Printf("the value of map is %v memoryLocation %p inside the function  \n", aMap, &aMap)
}

func addValueToMap1(hitsMap map[string]map[string]int, url string, country string) {
    countryMap, ok := hitsMap[url]
    if !ok {
        countryMap = make(map[string]int)
        hitsMap[url] = countryMap
    }
    hitsMap[url][country]++
	// hitsMap[url][country]++
	// alternative 2 to countryMap[country]++
	// hits, ok := countryMap[country]
	// if !ok {
	// 	countryMap[country] = 1
	// } else {
	// 	countryMap[country] = hits + 1
	// }

}

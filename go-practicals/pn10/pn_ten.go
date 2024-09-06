package pn10

import "fmt"

/*
To find frequency of each character of the string
for ex
input:chocolate
output:map[a:1 c:2 e:1 h:1 l:1 o:2 t:1]
*/

func CalculateFunction() {
	var inputString string
	fmt.Println("enter the string you want for frequency calculation")
	fmt.Scanln(&inputString)
	r := getCharFrequency(inputString)
	fmt.Println(r)
}

func getCharFrequency(gString string) map[string]int {
	freqMap := make(map[string]int)
	for _, v := range gString {
		freqMap[string(v)]++
	}
	return freqMap
}

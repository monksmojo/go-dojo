package pn17

import "fmt"

// given two string of same length find the distinct character  in string 2
// input string 1 "aabbcv"
// input string 2 "dcvaab"
// the character which is distinct in string2 is d
// input string 1 "aabbcv"
// input string 2 "aabbcc"
// the character which is distinct in string2 is c

// Find uncommon characters of the two strings

func FindUncommonCharacter() {
	st1 := "pork"
	st2 := "park"
	fetchUnCommonCharacter(st1, st2)

	st1 = "din"
	st2 = "fin"
	fetchUnCommonCharacter(st1, st2)
}

func fetchUnCommonCharacter(st1, st2 string) {
	st1Map := make(map[rune]int)
	for _, v := range st1 {
		st1Map[v] = st1Map[v] + 1
	}
	for _, v := range st2 {
		if st1Map[v] == 0 {
			fmt.Printf("%c\n", v)
		} else {
			st1Map[v] = st1Map[v] - 1
		}
	}
}

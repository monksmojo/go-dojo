package pn20

import (
	"fmt"
	"strings"
)

func searchWordInSentence() {
	sentence := "what you know about rolling down in the deep"
	word := "YOU"
	index := searchWord(getSliceOfWords, sentence, word)
	if index != -1 {
		fmt.Printf("found character %s in string %s at position %d \n", sentence, word, index+1)
	} else {
		fmt.Printf("%s character not found in string %s at position %d \n", sentence, word, index+1)
	}

}

func searchWord(getSliceOfWords func(string) []string, sentence, word string) int {
	for i, v := range getSliceOfWords(sentence) {
		if strings.ToLower(v) == strings.ToLower(word) {
			return i
		}
	}
	return -1
}

func getSliceOfWords(sentence string) []string {
	return strings.Split(sentence, " ")
}

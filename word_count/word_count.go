package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "count each word in in in in the the string"

	counter := WordCount(sampleString, " ")

	fmt.Println(counter)
}

func WordCount(stroke string, sep string) map[string]int {
	wordCount := make(map[string]int)

	slice := strings.Split(stroke, sep)
	fmt.Println(slice)

	for _, word := range slice {

		if doesExist(word, stroke) == true {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}

	}
	return wordCount

}

func doesExist(word string, stroke string) bool {
	return strings.Contains(stroke, word)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "count each word in in in in the the string"
	counter := WordCount(sampleString)

	for key, value := range counter {
		fmt.Printf("%v: %v\n", key, value)
	}
}

func WordCount(stroke string) map[string]int {
	slice := strings.Fields(stroke)
	wordCount := make(map[string]int)

	for _, word := range slice {
		wordCount[word]++
	}

	return wordCount
}

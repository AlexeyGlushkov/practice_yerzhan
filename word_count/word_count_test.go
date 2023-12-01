package main

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	//TODO: Add TableDriven implementation

	sampleText := "SOME some CASE caSE"

	expected := make(map[string]int, 0)
	expected["some"] = 2
	expected["case"] = 2

	result := WordCount(sampleText)

	//TODO: реализовать проверку результата, вне зависимости от регистра

}

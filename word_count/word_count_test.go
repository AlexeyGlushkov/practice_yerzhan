package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {

	tests := []struct {
		input string

		expected map[string]int
	}{

		{

			input: "count each word in in in in the the string",

			expected: map[string]int{

				"count": 1,

				"each": 1,

				"word": 1,

				"in": 4,

				"the": 2,

				"string": 1,
			},
		},

		{

			input: "",

			expected: map[string]int{},
		},
	}

	for _, test := range tests {

		result := WordCount(test.input)

		if !reflect.DeepEqual(result, test.expected) {

			t.Errorf("expected: %v, got: %v", test.expected, result)

		}

	}

}

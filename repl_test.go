package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "Hello World",
			output: []string{"hello", "world"},
		},
		{
			input:  "   string     with many    spaces   ",
			output: []string{"string", "with", "many", "spaces"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.output) {
			t.Errorf("Expected %v, got %v", c.output, actual)
		}

		for i := range actual {
			if actual[i] != c.output[i] {
				t.Errorf("Expected %v, got %v", c.output, actual)
			}
		}
	}
}

package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello worLd  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error: expected length of slice: %d, got=%d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("Error: expected word: %s, got=%s", expectedWord, word)
			}
		}
	}
}

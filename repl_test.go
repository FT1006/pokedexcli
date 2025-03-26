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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello     world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello! world ! ",
			expected: []string{"hello!", "world", "!"},
		},
		{
			input:    "hElLo! wOrLd !",
			expected: []string{"hello!", "world", "!"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Expected slice length %d, got %d", len(c.expected), len(actual))
		}
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("Expected word %s, got %s", expectedWord, word)
			}
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}

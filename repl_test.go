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
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    " a B c ",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of slice does not match")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not match expected word, %v", word, expectedWord)
			}
		}
	}
}

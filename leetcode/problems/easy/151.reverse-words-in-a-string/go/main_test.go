package main

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			input:    "a good   example",
			expected: "example good a",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			result := reverseWords(c.input)
			if result != c.expected {
				t.Errorf("expected %s, but got %s", c.expected, result)
			}
		})
	}
}

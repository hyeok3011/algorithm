package main

import "testing"

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "IceCreAm",
			expected: "AceCreIm",
		},
		{
			input:    "leetcode",
			expected: "leotcede",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			result := reverseVowels(c.input)
			if result != c.expected {
				t.Errorf("expected %s, but got %s", c.expected, result)
			}
		})
	}
}

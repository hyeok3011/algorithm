package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  string
	}{
		{
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
		{
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
	}

	for _, test := range tests {
		got := mergeAlternately(test.word1, test.word2)
		if got != test.want {
			t.Errorf("got %q, want %q", got, test.want)
		}
	}
}

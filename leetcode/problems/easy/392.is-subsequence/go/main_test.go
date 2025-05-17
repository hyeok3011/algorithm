package main

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, test := range tests {
		if got := isSubsequence(test.s, test.t); got != test.want {
			t.Errorf("isSubsequence(%q, %q) = %v, want %v", test.s, test.t, got, test.want)
		}
	}
}

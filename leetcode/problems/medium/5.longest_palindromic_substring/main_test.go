package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, longestPalindrome("aabaa"), "aabaa")
	// assert.Equal(t, longestPalindrome("cbbd"), "bb")
}

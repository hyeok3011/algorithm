package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestPalindrome(t *testing.T) {
	// assert.Equal(t, longestPalindrome("baab"), "baab")
	assert.Equal(t, longestPalindrome("babad"), "bb")
}

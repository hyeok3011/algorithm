package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bbbbb"), 1)
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)

}

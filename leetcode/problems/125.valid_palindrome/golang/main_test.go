package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, isPalindrome("A man, a plan, a canal: Panama"), true)
	assert.Equal(t, isPalindrome("race a car"), false)
	assert.Equal(t, isPalindrome(" "), true)

	assert.Equal(t, AnotherSolution("A man, a plan, a canal: Panama"), true)
	assert.Equal(t, AnotherSolution("race a car"), false)
	assert.Equal(t, AnotherSolution(" "), true)
}

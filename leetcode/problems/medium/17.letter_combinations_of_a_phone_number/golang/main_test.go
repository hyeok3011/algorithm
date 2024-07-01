package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLetterCombination(t *testing.T) {
	assert.Equal(t, letterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
	assert.Equal(t, letterCombinations(""), []string{})
	assert.Equal(t, letterCombinations("2"), []string{"a", "b", "c"})
}

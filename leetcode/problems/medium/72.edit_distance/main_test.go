package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinDistance(t *testing.T) {
	assert.Equal(t, minDistance("horse", "ros"), 3)
	assert.Equal(t, minDistance("intention", "execution"), 5)
	assert.Equal(t, minDistance("asd", "rasd"), 1)
}

// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')

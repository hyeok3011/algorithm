package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(3, 4, 2, 3, 3, 1, 5), "dllrl")
	assert.Equal(t, solution(2, 2, 1, 1, 2, 2, 2), "dr")
	assert.Equal(t, solution(3, 3, 1, 2, 3, 3, 4), "impossible")
	assert.Equal(t, solution(6, 6, 2, 6, 6, 5, 11), "ddddllllrrr")
	// , k
}

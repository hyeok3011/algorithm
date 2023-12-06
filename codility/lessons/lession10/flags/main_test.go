package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSol(t *testing.T) {
	assert.Equal(t, Solution([]int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}), 3)
	assert.Equal(t, Solution([]int{1, 2, 1}), 1)
}

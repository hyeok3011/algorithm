package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, Solution([]int{3, 2, 6, -1, 4, 5, -1, 2}), 17)
}

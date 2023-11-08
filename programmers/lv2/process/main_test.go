package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	// assert.Equal(t, solution([]int{2, 1, 3, 2}, 2), 1)
	// assert.Equal(t, solution([]int{1, 1, 9, 1, 1, 1}, 0), 5)
}

func TestOtherSolution(t *testing.T) {
	assert.Equal(t, otherSolution([]int{2, 1, 3, 2}, 2), 1)
	assert.Equal(t, otherSolution([]int{1, 1, 9, 1, 1, 1}, 0), 5)
}

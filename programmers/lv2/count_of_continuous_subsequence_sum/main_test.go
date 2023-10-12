package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]int{7, 9, 1, 1, 4}), 18)
	assert.Equal(t, solutionDP([]int{7, 9, 1, 1, 4}), 18)
}

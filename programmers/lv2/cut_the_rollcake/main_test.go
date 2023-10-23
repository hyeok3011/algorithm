package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]int{1, 2, 1, 3, 1, 4, 1, 2}), 2)
	assert.Equal(t, solution([]int{1, 2, 3, 1, 4}), 0)
}

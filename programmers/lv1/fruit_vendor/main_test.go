package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(3, 4, []int{1, 2, 3, 1, 2, 3, 1}), 8)
	assert.Equal(t, solution(4, 3, []int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2}), 33)
}

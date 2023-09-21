package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(8, 4, []int{2, 3, 6}), 2)
	assert.Equal(t, solution(5, 4, []int{1, 3}), 1)
	assert.Equal(t, solution(4, 1, []int{1, 2, 3, 4}), 4)
}

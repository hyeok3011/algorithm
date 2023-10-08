package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]int{3, 2, 7, 2}, []int{4, 6, 5, 1}), 2)
	assert.Equal(t, solution([]int{1, 2, 1, 2}, []int{1, 10, 1, 2}), 7)
	assert.Equal(t, solution([]int{1, 1}, []int{1, 5}), -1)
}

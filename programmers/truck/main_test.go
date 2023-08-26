package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	// assert.Equal(t, solution(5, 5, []int{1, 1, 1, 1, 1, 1}), 12)
	assert.Equal(t, solution(2, 10, []int{7, 4, 5, 6}), 8)
	assert.Equal(t, solution(100, 100, []int{10}), 101)
	assert.Equal(t, solution(100, 100, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}), 110)
	assert.Equal(t, solution(5, 5, []int{2, 2, 2, 2, 1, 1, 1, 1, 1}), 19)

}

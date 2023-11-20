package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	// assert.Equal(t, solution2([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, solution([]int{4, 1, 2, 1}, 4), 2)
	assert.Equal(t, solution([]int{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50}, 500), 15504)
	assert.Equal(t, solution([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 10), 15504)
	assert.Equal(t, solution([]int{25, 8, 17, 31, 12, 22, 14, 19, 11, 13}, 100), 5)

}

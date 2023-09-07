package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	input := [][]int{{4, 5}, {4, 8}, {10, 14}, {11, 13}, {5, 12}, {3, 7}, {1, 4}}
	// assert.Equal(t, solution(input), 3)
	assert.Equal(t, solution2(input), 3)
}

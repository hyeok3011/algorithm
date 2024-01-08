package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([][]int{
		{2, 3}, {4, 3}, {1, 1}, {2, 1},
	}), []int{2, 1, 1, 0})

	assert.ElementsMatch(t, solution([][]int{
		{4, 11}, {1, 12}, {8, 3}, {12, 7}, {4, 2}, {7, 11}, {4, 8}, {9, 6}, {10, 11}, {6, 10}, {3, 5}, {11, 1}, {5, 3}, {11, 9}, {3, 8},
	}), []int{4, 0, 1, 2})

	assert.ElementsMatch(t, solution([][]int{
		{4, 3}, {2, 3}, {1, 1}, {2, 1},
	}), []int{2, 1, 1, 0})
}

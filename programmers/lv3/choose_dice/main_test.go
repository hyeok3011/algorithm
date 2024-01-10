package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([][]int{
		{40, 41, 42, 43, 44, 45},
		{43, 43, 42, 42, 41, 41},
		{1, 1, 80, 80, 80, 80},
		{70, 70, 1, 1, 70, 70},
	}), []int{1, 3})

	assert.ElementsMatch(t, solution([][]int{
		{1, 2, 3, 4, 5, 6},
		{3, 3, 3, 3, 4, 4},
		{1, 3, 3, 4, 4, 4},
		{1, 1, 4, 4, 5, 5},
	}), []int{1, 4})
}

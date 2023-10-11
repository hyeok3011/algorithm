package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoltuion(t *testing.T) {
	inputData := [][]int{
		{2, 2, 5, 4}, {3, 3, 6, 6}, {5, 1, 6, 3},
	}
	assert.ElementsMatch(t, solution(6, 6, inputData), []int{8, 10, 25})

	inputData2 := [][]int{
		{1, 1, 2, 2}, {1, 2, 2, 3}, {2, 1, 3, 2}, {2, 2, 3, 3},
	}
	assert.ElementsMatch(t, solution(3, 3, inputData2), []int{1, 1, 5, 3})

	assert.ElementsMatch(t, solution(100, 97, [][]int{{1, 1, 100, 97}}), []int{1})
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	assert.ElementsMatch(t, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8), [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	})
	assert.ElementsMatch(t, combinationSum2([]int{2, 5, 2, 1, 2}, 5), [][]int{
		{1, 2, 2},
		{5},
	})
}

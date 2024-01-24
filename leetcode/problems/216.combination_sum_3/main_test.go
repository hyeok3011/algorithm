package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	assert.ElementsMatch(t, combinationSum3(3, 9), [][]int{
		{1, 2, 6}, {1, 3, 5}, {2, 3, 4},
	})
}

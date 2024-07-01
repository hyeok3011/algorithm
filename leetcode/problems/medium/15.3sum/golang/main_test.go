package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	assert.ElementsMatch(t, threeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	})

	assert.ElementsMatch(t, threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}), [][]int{
		{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1},
	})

}

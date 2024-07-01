package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	assert.ElementsMatch(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	assert.ElementsMatch(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	assert.ElementsMatch(t, searchRange([]int{}, 0), []int{-1, -1})
	assert.ElementsMatch(t, searchRange([]int{1}, 0), []int{-1, -1})
	// assert.ElementsMatch()
}

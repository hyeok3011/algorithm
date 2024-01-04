package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.ElementsMatch(t, twoSum([]int{2, 3, 4}, 6), []int{1, 3})
}

func TestOptimizedTwoSum(t *testing.T) {
	assert.ElementsMatch(t, optimizedTwoSum([]int{2, 3, 4}, 6), []int{1, 3})
}

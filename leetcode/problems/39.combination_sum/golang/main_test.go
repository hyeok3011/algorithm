package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	assert.ElementsMatch(t, combinationSum([]int{2, 3, 6, 7}, 7), [][]int{{2, 2, 3}, {7}})
}

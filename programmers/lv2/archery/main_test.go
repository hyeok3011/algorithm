package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution(5, []int{2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}), []int{0, 2, 2, 0, 1, 0, 0, 0, 0, 0, 0})
	// 10  7
	// 9 8 6
	assert.ElementsMatch(t, solution(1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), []int{-1})
	//
	assert.ElementsMatch(t, solution(9, []int{0, 0, 1, 2, 0, 1, 1, 1, 1, 1, 1}), []int{1, 1, 2, 0, 1, 2, 2, 0, 0, 0, 0})
	//
	assert.ElementsMatch(t, solution(10, []int{0, 0, 0, 0, 0, 0, 0, 0, 3, 4, 3}), []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 2})
}

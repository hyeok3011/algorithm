package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([]int{2, 3, 3, 5}), []int{3, 5, 5, -1})

	assert.ElementsMatch(t, solution([]int{9, 1, 5, 3, 6, 2}), []int{-1, 5, 6, 6, -1, -1})
}

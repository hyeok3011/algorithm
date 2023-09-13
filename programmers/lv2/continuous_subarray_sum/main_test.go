package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([]int{1, 2, 3, 4, 5}, 7), []int{2, 3})

	assert.ElementsMatch(t, solution([]int{1, 1, 1, 2, 3, 4, 5}, 5), []int{6, 6})
}

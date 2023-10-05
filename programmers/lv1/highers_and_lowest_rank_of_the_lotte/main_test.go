package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([]int{44, 1, 0, 0, 31, 25}, []int{31, 10, 45, 1, 6, 19}), []int{3, 5})
	assert.ElementsMatch(t, solution([]int{0, 0, 0, 0, 0, 0}, []int{38, 19, 20, 40, 15, 25}), []int{1, 6})
	assert.ElementsMatch(t, solution([]int{20, 9, 3, 45, 4, 35}, []int{20, 9, 3, 45, 4, 35}), []int{1, 1})

	assert.ElementsMatch(t, solution2([]int{44, 1, 0, 0, 31, 25}, []int{31, 10, 45, 1, 6, 19}), []int{3, 5})
	assert.ElementsMatch(t, solution2([]int{0, 0, 0, 0, 0, 0}, []int{38, 19, 20, 40, 15, 25}), []int{1, 6})
	assert.ElementsMatch(t, solution2([]int{20, 9, 3, 45, 4, 35}, []int{20, 9, 3, 45, 4, 35}), []int{1, 1})
}

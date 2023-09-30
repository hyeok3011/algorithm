package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([][]int{{40, 10000}, {25, 10000}}, []int{7000, 9000}), []int{1, 5400})
	assert.ElementsMatch(t, solution2([][]int{{40, 2900}, {23, 10000}, {11, 5200}, {5, 5900}, {40, 3100}, {27, 9200}, {32, 6900}}, []int{1300, 1500, 1600, 4900}), []int{4, 13860})
}

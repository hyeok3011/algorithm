package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution(2), [][]int{{1, 2}, {1, 3}, {2, 3}})
}

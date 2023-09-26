package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([]string{"X591X", "X1X5X", "X231X", "1XXX1"}), []int{1, 1, 27})
}

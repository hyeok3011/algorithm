package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	minerals := []string{"diamond", "diamond", "diamond", "iron", "iron", "diamond", "iron", "stone"}
	picks := []int{1, 3, 2}
	assert.Equal(t, solution(picks, minerals), 12)

	minerals2 := []string{"diamond", "diamond", "diamond", "diamond", "diamond", "stone", "stone", "stone", "stone", "stone", "stone", "stone", "stone"}
	picks2 := []int{0, 1, 1}
	assert.Equal(t, solution(picks2, minerals2), 30)

	assert.Equal(t, solution([]int{1, 1, 1}, []string{}), 0)
}

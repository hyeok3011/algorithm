package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]int{10, 17}, []int{5, 20}), 0)
	assert.Equal(t, solution([]int{10, 20}, []int{5, 17}), 10)
	assert.Equal(t, solution([]int{14, 35, 119}, []int{18, 30, 102}), 7)
}

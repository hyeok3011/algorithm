package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]int{6, 10, 2}), "6210")
	assert.Equal(t, solution([]int{3, 30, 34, 5, 9}), "9534330")
}

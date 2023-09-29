package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(4, 5, []int{1, 0, 3, 1, 2}, []int{0, 3, 0, 4, 0}), int64(16))
}

package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]int{2, 3, -6, 1, 3, -1, 2, 4}), int64(10))
}

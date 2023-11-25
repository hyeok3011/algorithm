package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(8, 4, 7), 3)
	assert.Equal(t, solution(4, 1, 2), 1)
	assert.Equal(t, solution(12, 1, 10), 4)
}

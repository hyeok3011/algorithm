package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(10), 3)
	assert.Equal(t, solution(12), 11)
}

package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(4), int64(5))
	assert.Equal(t, solution(3), int64(3))
}

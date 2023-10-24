package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(16), 6)
	assert.Equal(t, solution(2554), 16)
}

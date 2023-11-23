package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution("JEROEN"), 56)
	assert.Equal(t, solution("JAN"), 23)
	assert.Equal(t, solution("ABABAAAAABA"), 10)
}

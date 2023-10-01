package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution("3141592", "271"), 2)
	assert.Equal(t, solution("10203", "15"), 3)
}

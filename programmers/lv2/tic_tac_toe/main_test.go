package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]string{"O.X", ".O.", "..X"}), 1)
	assert.Equal(t, solution([]string{"OOO", "...", "XXX"}), 0)
	assert.Equal(t, solution([]string{"...", ".X.", "..."}), 0)
	assert.Equal(t, solution([]string{"...", "...", "..."}), 1)
}

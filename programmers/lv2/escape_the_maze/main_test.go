package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]string{"SOOOL", "XXXXO", "OOOOO", "OXXXX", "OOOOE"}), 16)
	assert.Equal(t, solution([]string{"LOOXS", "OOOOX", "OOOOO", "OOOOO", "EOOOO"}), -1)
}

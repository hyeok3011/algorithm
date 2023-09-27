package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]string{"i", "drink", "water"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"}), "Yes")
	assert.Equal(t, solution([]string{"i", "water", "drink"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"}), "No")
}

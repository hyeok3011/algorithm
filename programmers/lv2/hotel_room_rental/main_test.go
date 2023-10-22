package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([][]string{
		{"15:00", "17:00"}, {"16:40", "18:20"}, {"14:20", "15:20"}, {"14:10", "19:20"}, {"18:20", "21:20"},
	}), 3)
	assert.Equal(t, solution([][]string{
		{"09:10", "10:10"}, {"10:20", "12:20"},
	}), 1)
	assert.Equal(t, solution([][]string{
		{"10:20", "12:30"}, {"10:20", "12:30"}, {"10:20", "12:30"},
	}), 3)
}

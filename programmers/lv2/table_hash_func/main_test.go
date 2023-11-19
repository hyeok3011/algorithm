package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([][]int{{2, 2, 6}, {1, 5, 10}, {4, 2, 9}, {3, 8, 3}}, 2, 2, 3), 4)
}

// {4, 2, 9}
// {2, 2, 6}
// {1, 5, 10}
// {3, 8, 3}

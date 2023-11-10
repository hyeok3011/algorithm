package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(8, [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {5, 6}, {5, 7}, {5, 8}}), 2)
	assert.Equal(t, solution(10, [][]int{{4, 1}, {5, 1}, {5, 6}, {7, 6}, {1, 2}, {1, 3}, {6, 8}, {2, 9}, {9, 10}}), 3)
	assert.Equal(t, solution(8, [][]int{{1, 2}, {2, 5}, {1, 4}, {4, 6}, {1, 3}, {3, 7}, {7, 8}}), 4)
	assert.Equal(t, solution(6, [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {3, 6}}), 2)
	assert.Equal(t, solution(6, [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}, {5, 6}}), 3)
	assert.Equal(t, solution(4, [][]int{{1, 2}, {2, 3}, {3, 4}}), 2)
}

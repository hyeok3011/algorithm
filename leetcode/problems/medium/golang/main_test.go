package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinPathSum(t *testing.T) {
	assert.Equal(t, minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}), 7)

	assert.Equal(t, minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}), 12)
}

// 1,2,5,
// 3,2,1

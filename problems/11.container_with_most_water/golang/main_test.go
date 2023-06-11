package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	assert.Equal(t, maxArea([]int{1, 1}), 1)
}

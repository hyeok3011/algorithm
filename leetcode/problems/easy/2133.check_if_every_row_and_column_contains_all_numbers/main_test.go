package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCheckValid(t *testing.T) {
	assert.Equal(t, checkValid([][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}), true)
	assert.Equal(t, checkValid([][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}), false)
}

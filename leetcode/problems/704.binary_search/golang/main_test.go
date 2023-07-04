package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, search([]int{-1, 0, 3, 5, 9, 12}, 9), 4)
	assert.Equal(t, search([]int{-1, 0, 3, 5, 9, 12}, 2), -1)
	assert.Equal(t, search([]int{5}, 5), 0)
}

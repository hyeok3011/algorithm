package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	assert.Equal(t, search([]int{6, 7, 8, 9, 10, 11, 0, 1}, 12), -1)
	assert.Equal(t, search([]int{6, 7, 8, 9, 10, 11, 0, 1, 2, 4, 5}, 7), 1)
	assert.Equal(t, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8), 4)
	assert.Equal(t, search([]int{3, 1}, 1), 1)
}

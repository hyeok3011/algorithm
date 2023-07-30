package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCanJump(t *testing.T) {
	assert.Equal(t, canJump([]int{2, 3, 1, 1, 4}), true)
	assert.Equal(t, canJump([]int{3, 2, 1, 0, 4}), false)
	assert.Equal(t, canJump([]int{2, 0}), true)
}

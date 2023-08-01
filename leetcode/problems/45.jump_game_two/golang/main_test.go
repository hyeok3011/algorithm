package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestJump(t *testing.T) {
	// assert.Equal(t, jump([]int{2, 3, 0, 1, 4}), 2)
	// assert.Equal(t, jump([]int{2, 3, 1, 1, 4}), 2)
	// assert.Equal(t, jump([]int{2, 3, 1}), 1)
	// assert.Equal(t, jump([]int{1, 2, 3, 4, 5}), 3)
	// assert.Equal(t, jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}), 2)
	assert.Equal(t, jump([]int{5, 4, 0, 1, 3, 6, 8, 0, 9, 4, 9, 1, 8, 7, 4, 8}), 3)

}

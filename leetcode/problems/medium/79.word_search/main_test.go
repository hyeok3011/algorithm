package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestExist(t *testing.T) {
	assert.Equal(t, exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"), true)

	assert.Equal(t, exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB"), false)

	assert.Equal(t, exist([][]byte{
		{'a'},
		{'b'},
	}, "ba"), true)
}

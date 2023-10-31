package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFormingMagicSquare(t *testing.T) {
	assert.Equal(t, formingMagicSquare([][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}), int32(7))
}

// {{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}

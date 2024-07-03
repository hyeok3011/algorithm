package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsPerfectSquare(t *testing.T) {
	assert.Equal(t, isPerfectSquare(100), true)
	assert.Equal(t, isPerfectSquare(16), true)
	assert.Equal(t, isPerfectSquare(14), false)
}

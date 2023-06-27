package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSqrt(t *testing.T) {
	assert.Equal(t, mySqrt(0), 0)
	assert.Equal(t, mySqrt(2), 1)
	assert.Equal(t, mySqrt(3), 1)
	assert.Equal(t, mySqrt(4), 2)
	assert.Equal(t, mySqrt(8), 2)
	assert.Equal(t, mySqrt(10), 3)
}

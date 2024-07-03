package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFib(t *testing.T) {
	assert.Equal(t, fib(2), 1)
	assert.Equal(t, fib(3), 2)
	assert.Equal(t, fib(4), 3)
}

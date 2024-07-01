package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, divide(10, 3), 3)
	assert.Equal(t, divide(7, -3), -2)
}

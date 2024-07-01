package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestUniquePaths(t *testing.T) {
	// assert.Equal(t, uniquePaths(3, 2), 3)
	assert.Equal(t, uniquePaths(3, 7), 28)
}

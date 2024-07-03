package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsHappy(t *testing.T) {
	assert.Equal(t, isHappy(19), true)
	assert.Equal(t, isHappy(2), false)
}

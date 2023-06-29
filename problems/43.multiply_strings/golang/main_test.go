package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, multiply("123", "456"), "56088")
	// assert.Equal(t, multiply("123", "456"), "56088")
	assert.Equal(t, multiply("2", "3"), "6")
	assert.Equal(t, multiply("498828660196", "840477629533"), "419254329864656431168468")
}

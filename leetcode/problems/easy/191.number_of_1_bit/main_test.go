package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, hammingWeight(00000000000000000000000000001011), 3)
}

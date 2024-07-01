package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, convert("", 3), 3)
}

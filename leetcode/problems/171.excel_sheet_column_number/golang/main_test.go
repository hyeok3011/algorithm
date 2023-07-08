package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTitleToNumber(t *testing.T) {
	assert.Equal(t, titleToNumber("A"), 1)
	assert.Equal(t, titleToNumber("AB"), 28)
	assert.Equal(t, titleToNumber("ZY"), 701)
}

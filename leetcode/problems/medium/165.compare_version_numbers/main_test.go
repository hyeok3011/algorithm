package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCompareVersion(t *testing.T) {
	assert.Equal(t, compareVersion("1.01", "1.001"), 0)
	assert.Equal(t, compareVersion("1.0", "1.0.0"), 0)
	assert.Equal(t, compareVersion("0.1", "1.1"), -1)
}

package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, isSubsequence("abc", "ahbgdc"), true)
}

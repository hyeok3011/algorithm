package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsSameTree(t *testing.T) {
	assert.Equal(t, isSameTree(nil, nil), true)
}

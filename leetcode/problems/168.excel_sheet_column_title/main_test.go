package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConvertToTitle(t *testing.T) {
	assert.Equal(t, convertToTitle(1), "A")
	assert.Equal(t, convertToTitle(28), "AB")
	assert.Equal(t, convertToTitle(701), "ZY")
	assert.Equal(t, convertToTitle(52), "AZ")
}

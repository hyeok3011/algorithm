package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAddBinary(t *testing.T) {
	assert.Equal(t, addBinary("11", "1"), "100")
	assert.Equal(t, addBinary("1010", "1011"), "10101")
	assert.Equal(t, addBinary("1111", "1111"), "11110")
}

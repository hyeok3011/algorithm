package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, intToRoman(3), "III")
	assert.Equal(t, intToRoman(58), "LVIII")
	assert.Equal(t, intToRoman(1994), "MCMXCIV")
}

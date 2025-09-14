package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, myAtoi("9223372036854775808"), -2147483648)
}

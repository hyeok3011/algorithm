package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCountAndSay(t *testing.T) {
	// assert.Equal(t, countAndSay(1), "1")
	assert.Equal(t, countAndSay(4), "1211")
	assert.Equal(t, countAndSay(5), "111221")
	assert.Equal(t, countAndSay(10), "13211311123113112211")
}

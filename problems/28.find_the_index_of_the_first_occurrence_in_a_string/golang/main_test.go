package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, strStr("sadbutsad", "sad"), 0)
	assert.Equal(t, strStr("leetcode", "leeto"), -1)
}

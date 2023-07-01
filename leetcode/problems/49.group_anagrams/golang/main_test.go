package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGroupAnagrams(t *testing.T) {
	assert.Equal(t, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	})
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	assert.ElementsMatch(t, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	})
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitIntoFibonacci(t *testing.T) {
	assert.ElementsMatch(t, splitIntoFibonacci("1101111"), []int{11, 0, 11, 11})
	// assert.ElementsMatch(t, splitIntoFibonacci("112358130"), []int{})
	// assert.ElementsMatch(t, splitIntoFibonacci("232252752"), []int{23, 2, 25, 27, 52})
}

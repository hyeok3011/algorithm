package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNonDivisibleSubset(t *testing.T) {
	assert.Equal(t, nonDivisibleSubset(4, []int32{19, 10, 12, 10, 24, 25, 22}), int32(3))
	assert.Equal(t, nonDivisibleSubset(3, []int32{1, 7, 4, 2}), int32(3))
	assert.Equal(t, nonDivisibleSubset(7, []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}), int32(11))
}

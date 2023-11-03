package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHourglassSum(t *testing.T) {
	input := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	assert.Equal(t, hourglassSum(input), int32(19))
}

package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, singleNumber([]int{1, 2, 2}), 1)
	assert.Equal(t, singleNumber([]int{4, 2, 2, 1, 1}), 4)
}

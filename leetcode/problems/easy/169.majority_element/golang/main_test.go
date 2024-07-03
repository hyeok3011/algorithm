package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMajorityElement(t *testing.T) {
	assert.Equal(t, majorityElement([]int{3, 2, 3}), 3)
	assert.Equal(t, anotherSolution([]int{3, 2, 3}), 3)
}

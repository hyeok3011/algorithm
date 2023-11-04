package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution("banana"), []int{-1, -1, -1, 2, 2, 2})
	assert.Equal(t, solution("foobar"), []int{-1, -1, 1, -1, -1, -1})
}

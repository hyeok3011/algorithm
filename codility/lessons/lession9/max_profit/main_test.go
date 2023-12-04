package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolition(t *testing.T) {
	assert.Equal(t, Solution([]int{23171, 21011, 21123, 21366, 21013, 21367}), 356)
}

package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution("1924", 2), "94")
	assert.Equal(t, solution("1231234", 3), "3234")
	assert.Equal(t, solution("4177252841", 4), "775841")
	assert.Equal(t, solution("7777", 2), "77")
}

// 76778 2

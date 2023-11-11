package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]string{"banana", "apple", "rice", "pork", "pot"}, []int{3, 2, 2, 2, 1}, []string{"chicken", "apple", "apple", "banana", "rice", "apple", "pork", "banana", "pork", "rice", "pot", "banana", "apple", "banana"}), 3)
	assert.Equal(t, solution([]string{"apple"}, []int{10}, []string{"banana", "banana", "banana", "banana", "banana", "banana", "banana", "banana", "banana", "banana"}), 0)
}

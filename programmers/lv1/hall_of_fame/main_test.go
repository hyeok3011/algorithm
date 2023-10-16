package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution(3, []int{10, 100, 20, 150, 1, 100, 200}), []int{10, 10, 10, 20, 20, 100, 100})
}

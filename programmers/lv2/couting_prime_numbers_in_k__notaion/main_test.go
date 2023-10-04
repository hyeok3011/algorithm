package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution(437674, 3), 3)
	assert.Equal(t, solution(110011, 10), 2)
}

func TestIsPrime(t *testing.T) {
	assert.Equal(t, isPrime(11), true)
	assert.Equal(t, isPrime(2), true)
	assert.Equal(t, isPrime(3), true)
	assert.Equal(t, isPrime(4), false)
}

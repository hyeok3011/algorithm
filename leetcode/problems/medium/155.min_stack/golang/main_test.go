package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	// ["MinStack","push","push","push","getMin","pop","top","getMin"]
	// [[],[-2],[0],[-3],[],[],[],[]]

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	assert.Equal(t, minStack.GetMin(), -3)
	minStack.Pop()
	assert.Equal(t, minStack.Top(), 0)
	assert.Equal(t, minStack.GetMin(), -2)
}

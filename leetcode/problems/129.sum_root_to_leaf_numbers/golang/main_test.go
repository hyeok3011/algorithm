package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSumNumbers(t *testing.T) {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
			Val: 9,
		},
		Right: &TreeNode{Val: 0},
		Val:   4,
	}
	assert.Equal(t, sumNumbers(input), 1026)
}

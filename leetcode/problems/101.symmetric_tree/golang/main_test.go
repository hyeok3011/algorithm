package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsSymmetricTree(t *testing.T) {
	input := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
			Val: 2,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
			Val: 2,
		},
		Val: 1,
	}
	assert.Equal(t, isSymmetricQueue(input), true)
	assert.Equal(t, isSymmetricRecursion(input), true)
}

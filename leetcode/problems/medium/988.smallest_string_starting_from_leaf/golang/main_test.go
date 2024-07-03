package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSmallestFromLeaf(t *testing.T) {
	input := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	assert.Equal(t, smallestFromLeaf(input), "dba")
}

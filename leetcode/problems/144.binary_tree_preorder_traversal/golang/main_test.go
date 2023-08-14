package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	inputNode := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	assert.ElementsMatch(t, preorderTraversal(inputNode), []int{1, 2, 3})
	assert.ElementsMatch(t, preorderTraversalRecursion(inputNode), []int{1, 2, 3})
}

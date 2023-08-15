package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	input := &TreeNode{
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
		Val: 3,
	}
	assert.ElementsMatch(t, postorderTraversal(input), []int{1, 2, 3})
	assert.ElementsMatch(t, postorderTraversalStack2(input), []int{1, 2, 3})
	assert.ElementsMatch(t, postorderTraversalRecursion(input), []int{1, 2, 3})
}

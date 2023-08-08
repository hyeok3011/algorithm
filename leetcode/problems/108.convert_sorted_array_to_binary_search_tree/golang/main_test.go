package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayToBST(t *testing.T) {
	assert.ElementsMatch(t, inorderTraversal(sortedArrayToBST([]int{-10, -3, 0, 5, 9})), []int{0, -3, 9, -10, 5})
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

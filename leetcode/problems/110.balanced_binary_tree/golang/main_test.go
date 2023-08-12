package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4

func TestIsBalanced(t *testing.T) {

	input := []int{1, 2, 2, 3, 0, 0, 3, 4, 0, 0, 4}
	assert.Equal(t, isBalanced(buildBST(0, len(input)-1, input)), true)
}

func buildBST(startIndex, endIndex int, nums []int) *TreeNode {
	if startIndex > endIndex {
		return nil
	}

	mid := startIndex + (endIndex-startIndex)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBST(startIndex, mid-1, nums)
	root.Right = buildBST(mid+1, endIndex, nums)
	return root
}

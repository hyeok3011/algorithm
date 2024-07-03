package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 0 ms Beats 100% Memory 3.5 MB Beats 61.24%
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(0, len(nums)-1, nums)
}

func buildBST(startIndex, endIndex int, nums []int) *TreeNode {
	if startIndex == endIndex {
		return &TreeNode{Val: nums[startIndex]}
	}
	mid := (endIndex + startIndex) / 2
	node := &TreeNode{Val: nums[mid]}

	if mid > startIndex {
		node.Left = buildBST(startIndex, mid-1, nums)
	}

	node.Right = buildBST(mid+1, endIndex, nums)

	return node
}

func buildBST2(startIndex, endIndex int, nums []int) *TreeNode {
	if startIndex > endIndex {
		return nil
	}

	mid := startIndex + (endIndex-startIndex)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBST(startIndex, mid-1, nums)
	root.Right = buildBST(mid+1, endIndex, nums)
	return root
}

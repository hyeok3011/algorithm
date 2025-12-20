package main

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	return recursion(root, -10001)
}

func recursion(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	count := 0
	if maxVal <= node.Val {
		count++
		maxVal = node.Val
	}

	count += recursion(node.Left, maxVal)
	count += recursion(node.Right, maxVal)
	return count
}

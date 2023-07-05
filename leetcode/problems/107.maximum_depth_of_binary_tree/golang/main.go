package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//	https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
//
// Runtime 0 ms Beats 100% Memory 4.3 MB Beats 42.25%
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 1)
}

func dfs(node *TreeNode, depth int) int {
	if node.Left == nil && node.Right == nil {
		return depth
	}

	leftDepth, rightDepth := 0, 0
	if node.Left != nil {
		leftDepth = dfs(node.Left, depth+1)
	}
	if node.Right != nil {
		rightDepth = dfs(node.Right, depth+1)
	}

	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	var curr int

	if left > right {
		curr = left
	} else {
		curr = right
	}

	return curr + 1
}

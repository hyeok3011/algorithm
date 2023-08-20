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

// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 64.86%
// https://leetcode.com/problems/sum-root-to-leaf-numbers/description/

func sumNumbers(root *TreeNode) int {
	return dfs2(root, 0)
}
func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	currentSum := sum*10 + node.Val
	newSum := 0
	newSum += dfs(node.Left, currentSum)
	newSum += dfs(node.Right, currentSum)
	if newSum > currentSum {
		return newSum
	}

	return currentSum
}

func dfs2(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	currentSum := sum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	return dfs2(node.Left, currentSum) + dfs2(node.Right, currentSum)
}

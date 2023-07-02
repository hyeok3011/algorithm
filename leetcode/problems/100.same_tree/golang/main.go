package main

// https://leetcode.com/problems/same-tree/description/
// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 80.72%

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || q.Val != p.Val {
		return false
	}

	return isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)
}

package main

import "math"

// https://leetcode.com/problems/validate-binary-search-tree/description/
func isValidBST(root *TreeNode) bool {
	return validBst(root, math.MinInt, math.MaxInt)
}

func validBst(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if min < node.Val && node.Val < max {
		return validBst(node.Left, min, node.Val) && validBst(node.Right, node.Val, max)
	}

	return false
}

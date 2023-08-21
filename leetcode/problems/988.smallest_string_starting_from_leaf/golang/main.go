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

// https://leetcode.com/problems/smallest-string-starting-from-leaf/description/
// Runtime 0 ms Beats 100% Memory 4.5 MB Beats 100%
// 사전순을 헷갈려서 조금 헤맸음..
func smallestFromLeaf(root *TreeNode) string {
	return dfs(root, "")
}

func dfs(node *TreeNode, suffix string) string {
	if node == nil {
		return ""
	}

	currentLetter := string(byte(node.Val) + byte('a'))
	if node.Left == nil {
		return dfs(node.Right, currentLetter) + currentLetter
	}

	if node.Right == nil {
		return dfs(node.Left, currentLetter) + currentLetter
	}

	leftPath := dfs(node.Left, currentLetter) + currentLetter
	rightPath := dfs(node.Right, currentLetter) + currentLetter

	if leftPath+suffix > rightPath+suffix {
		return rightPath
	}
	return leftPath
}

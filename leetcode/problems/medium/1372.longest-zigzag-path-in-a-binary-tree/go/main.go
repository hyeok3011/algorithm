package main

import (
	. "github.com/hyeok3011/algorithm/structures"
)

// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/?envType=study-plan-v2&envId=leetcode-75
// topdown방식
func longestZigZag(root *TreeNode) int {
	maxDepth := 0
	var recursion func(node *TreeNode, depth int, isRight bool)
	recursion = func(node *TreeNode, depth int, isRight bool) {
		if node == nil {
			return
		}

		maxDepth = max(maxDepth, depth)

		if isRight {
			recursion(node.Left, depth+1, false)
			recursion(node.Right, 1, true)
		} else {
			recursion(node.Right, depth+1, true)
			recursion(node.Left, 1, false)
		}

	}

	recursion(root, 0, false)
	return maxDepth
}

// 다른 사람의 TopDown방식
func longestZigZagBottomUp(root *TreeNode) int {
	maxLen := 0
	dfs(root, &maxLen)
	return maxLen
}

func dfs(node *TreeNode, maxLen *int) (goLeft, goRight int) {
	if node == nil {
		return 0, 0
	}

	_, leftRight := dfs(node.Left, maxLen)
	rightLeft, _ := dfs(node.Right, maxLen)

	if node.Left != nil {
		goLeft = 1 + leftRight
	} else {
		goLeft = 0
	}

	if node.Right != nil {
		goRight = 1 + rightLeft
	} else {
		goRight = 0
	}

	if goLeft > *maxLen {
		*maxLen = goLeft
	}
	if goRight > *maxLen {
		*maxLen = goRight
	}

	return goLeft, goRight
}

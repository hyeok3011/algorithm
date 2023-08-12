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

// https://leetcode.com/problems/balanced-binary-tree/description/
// Runtime 4 ms Beats 82.57% Memory 5.9 MB Beats 99.34%
// 처음에 height-balanced를 잘못이해하고 풀어서 애먹었음. 잊지말자 height-balanced
// height-balanced
// Height-balanced 이진 트리란, 각 노드의 서브트리의 높이 차이가 한 칸 이상을 초과하지 않는 이진 트리를 의미

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalancedBacktracking(root) != -1
}

func isBalancedBacktracking(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := isBalancedBacktracking(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := isBalancedBacktracking(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return a * -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

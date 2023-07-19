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

// https://leetcode.com/problems/path-sum/description/
// Runtime 4 ms Beats 82.10% Memory 4.6 MB Beats 98.89%
func hasPathSumDFSRecersion(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(node *TreeNode, totalSum, targetSum int) bool {
	if node == nil {
		return false
	}

	totalSum += node.Val

	if totalSum == targetSum {
		if node.Left == nil && node.Right == nil {
			return true
		}
	}

	if dfs(node.Left, totalSum, targetSum) {
		return true
	}

	if dfs(node.Right, totalSum, targetSum) {
		return true
	}

	return false
}

// Runtime 0 ms Beats 100% Memory 4.5 MB Beats 98.89%
func hasPathSumDFSStack(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	nodeStack := []*TreeNode{root}
	sumStack := []int{0}

	for len(nodeStack) == 0 {
		currentNode := nodeStack[len(nodeStack)-1]
		currentSum := sumStack[len(sumStack)-1]

		nodeStack = nodeStack[:len(nodeStack)-1]
		sumStack = sumStack[:len(sumStack)-1]

		currentSum += currentNode.Val
		if currentSum == targetSum && currentNode.Left == nil && currentNode.Right == nil {
			return true
		}

		if currentNode.Right != nil {
			nodeStack = append(nodeStack, currentNode.Right)
			sumStack = append(sumStack, currentSum)
		}

		if currentNode.Left != nil {
			nodeStack = append(nodeStack, currentNode.Left)
			sumStack = append(sumStack, currentSum)
		}
	}

	return false
}

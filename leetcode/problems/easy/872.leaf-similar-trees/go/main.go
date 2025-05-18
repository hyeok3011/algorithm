package main

// https://leetcode.com/problems/leaf-similar-trees/?envType=study-plan-v2&envId=leetcode-75

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

func leafSimilar(firstNode *TreeNode, secondNode *TreeNode) bool {
	firstNodeEndValues := []int{}
	dfs(firstNode, &firstNodeEndValues)
	secondNodeEndValues := []int{}
	dfs(secondNode, &secondNodeEndValues)

	if len(firstNodeEndValues) != len(secondNodeEndValues) {
		return false
	}

	for i := range firstNodeEndValues {
		if firstNodeEndValues[i] != secondNodeEndValues[i] {
			return false
		}
	}

	return true
}

func dfs(node *TreeNode, endValues *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*endValues = append(*endValues, node.Val)
	}
	dfs(node.Left, endValues)
	dfs(node.Right, endValues)
}

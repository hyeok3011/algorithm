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

func searchBST(root *TreeNode, val int) *TreeNode {
	currentNode := root
	for currentNode != nil {
		if currentNode.Val == val {
			break
		} else if currentNode.Val >= val {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	return currentNode
}

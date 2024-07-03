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

// https://leetcode.com/problems/binary-tree-preorder-traversal/description/
// Runtime 0 ms Beats 100% Memory 2 MB Beats 100%
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nodeStack := []*TreeNode{root}
	answer := []int{}

	for len(nodeStack) > 0 {
		currentNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		answer = append(answer, currentNode.Val)
		if currentNode.Right != nil {
			nodeStack = append(nodeStack, currentNode.Right)
		}

		if currentNode.Left != nil {
			nodeStack = append(nodeStack, currentNode.Left)
		}

	}
	return answer
}

func preorderTraversalRecursion(root *TreeNode) []int {
	values := []int{}
	preorderRecursion(root, &values)
	return values
}
func preorderRecursion(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	*(values) = append(*(values), node.Val)
	preorderRecursion(node.Left, values)
	preorderRecursion(node.Right, values)
}

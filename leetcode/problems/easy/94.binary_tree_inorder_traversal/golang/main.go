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

// https://leetcode.com/problems/binary-tree-inorder-traversal/description/
// Runtime 0 ms Beats 100% Memory 2 MB Beats 23.40%
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	path := []int{}
	path = append(path, inorderTraversal(root.Left)...)
	path = append(path, root.Val)
	path = append(path, inorderTraversal(root.Right)...)
	return path
}

func inorderTraversal2(root *TreeNode) []int {
	paths := []int{}
	inorderTraversalDfs(root, &paths)
	return paths
}

func inorderTraversalDfs(node *TreeNode, paths *[]int) {
	if node == nil {
		return
	}
	inorderTraversalDfs(node.Left, paths)
	*(paths) = append(*(paths), node.Val)
	inorderTraversalDfs(node.Right, paths)
}

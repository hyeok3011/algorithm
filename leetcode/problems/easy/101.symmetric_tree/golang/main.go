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

// Runtime 0 ms Beats 100% Memory 2.9 MB Beats 98.36%
// DFS recursion Version
func isSymmetricRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recursion(root.Left, root.Right)
}

func recursion(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	} else if (leftNode == nil || rightNode == nil) || (leftNode.Val != rightNode.Val) {
		return false
	}

	if !recursion(leftNode.Left, rightNode.Right) || !recursion(leftNode.Right, rightNode.Left) {
		return false
	}

	return true
}

// Runtime 0 ms Beats 100% Memory 2.9 MB Beats 17.76%
// BFS Queue Version
func isSymmetricQueue(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodeQueue := []*TreeNode{root.Left, root.Right}

	for len(nodeQueue) > 0 {
		leftsideNode := nodeQueue[0]
		rightsideNode := nodeQueue[1]

		nodeQueue = nodeQueue[2:]

		if leftsideNode == nil && rightsideNode == nil {
			continue
		} else if (leftsideNode == nil || rightsideNode == nil) || (leftsideNode.Val != rightsideNode.Val) {
			return false
		}

		nodeQueue = append(nodeQueue, leftsideNode.Left, rightsideNode.Right, leftsideNode.Right, rightsideNode.Left)
	}

	return true
}

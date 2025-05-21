package main

// https://leetcode.com/problems/binary-tree-right-side-view/?envType=study-plan-v2&envId=leetcode-75
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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	answer := []int{}

	for len(queue) > 0 {
		breadth := len(queue)
		answer = append(answer, queue[0].Val)
		for i := 0; i < breadth; i++ {
			currentNode := queue[i]
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
		}

		queue = queue[breadth:]
	}

	return answer
}

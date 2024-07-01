package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
// Runtime 185 ms Beats 96.40% Memory 17.8 MB Beats 77.7%
// 이미 최적화가 된거같은데 점수가 왔다갔다함.... 뭐가 문제인지 모르겠음...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		depth += 1
		queueLen := len(queue)

		for i := 0; i < queueLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				queue = []*TreeNode{}
				break
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}

	return depth
}

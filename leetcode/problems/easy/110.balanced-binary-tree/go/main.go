package main

// https://leetcode.com/problems/balanced-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    var recursion func(node *TreeNode) int 
    recursion = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftHeight := recursion(node.Left)
        if leftHeight == -1 {
            return -1
        }

        rightHeight := recursion(node.Right)
        if rightHeight == -1 {
            return -1
        }

        if abs(leftHeight - rightHeight) > 1 {
            return -1
        }

        return max(leftHeight, rightHeight) + 1
    }

    return recursion(root) != -1
}

func abs(v int) int {
    if v > 0 {
        return v
    }
    return v * -1
}

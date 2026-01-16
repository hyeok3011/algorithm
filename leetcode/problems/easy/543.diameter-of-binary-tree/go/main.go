package main

// https://leetcode.com/problems/diameter-of-binary-tree/?envType=problem-list-v2&envId=rab78cw1
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    var recursion func(node *TreeNode) int
    recursion = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftDepth := recursion(node.Left)
        rightDepth := recursion(node.Right)
        diameter = max(diameter, leftDepth + rightDepth)
        return max(leftDepth, rightDepth) + 1
    }
    recursion(root)
    return diameter
}

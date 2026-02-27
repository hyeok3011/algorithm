package main

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue := []*TreeNode{root}
    answer := [][]int{}
    reverse := false
    for len(queue) > 0 {
        levelSize := len(queue)
        values := make([]int, levelSize)
        for i := range levelSize {
            node := queue[i]
            if reverse {
                values[levelSize - 1 - i] = node.Val
            } else {
                values[i] = node.Val
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[levelSize:]
        reverse = !reverse
        answer = append(answer, values)
    }

    return answer
}

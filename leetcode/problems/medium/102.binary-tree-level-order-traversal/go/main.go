package main

//  https://leetcode.com/problems/binary-tree-level-order-traversal/?envType=problem-list-v2&envId=rab78cw1
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    result := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        levelSize := len(queue)
        values := make([]int, levelSize)
        for i := 0; i < levelSize; i++ {
            node := queue[i]
            values[i] = node.Val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        result = append(result, values)
        queue = queue[levelSize:]
    }

    return result
}

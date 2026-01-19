package main

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/?envType=problem-list-v2&envId=rab78cw1
func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    current := root
    for current != nil || len(stack) > 0 {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }

        current = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        k--
        if k == 0 {
            return current.Val
        }

        current = current.Right
    }

    return -1
}

// node지우는 방법, value swap까지는 적용하지 않음.
func kthSmallestDeleteNode(root *TreeNode, k int) int {

    var deleteSmallest func(node *TreeNode) *TreeNode
    deleteSmallest = func(node *TreeNode) *TreeNode {
        if node.Left == nil {
            return node.Right
        }

        node.Left = deleteSmallest(node.Left)
        return node
    }

    for i := 0; i < k - 1; i++ {
        root = deleteSmallest(root)
    }

    return findMin(root).Val
}

func findMin(node *TreeNode) *TreeNode {
    if node.Left == nil {
        return node
    }
    return findMin(node.Left)
}

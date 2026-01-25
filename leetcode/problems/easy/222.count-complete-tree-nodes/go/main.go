package main

// https://leetcode.com/problems/count-complete-tree-nodes/
// @@@@ complete tree를 제대로 이해하지 못하고 있었음. 좋은 문제같음.
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHeight := getHeight(root.Left)
    rightHeight := getHeight(root.Right)
    if leftHeight == rightHeight {
        return ((1 << leftHeight) - 1) + countNodes(root.Right) + 1
    }
    return countNodes(root.Left) + ((1 << rightHeight) - 1) + 1
}

func getHeight(node *TreeNode) int {
    height := 0
    for node != nil {
        node = node.Left
        height++
    }
    return height
}

// 다른 사람의 풀이, 높이를 왼쪽 오른쏙으로 각각탐색해서 완벽한 삼각형인지를 확인함, 생각하지 못한 접근 방법.
func AnotherSolution(root *TreeNode) int {
    left := leftHeight(root)
    right := rightHeight(root)

    if left == right {
        return int(math.Pow(2, float64(left))) - 1
    } else {
        return countNodes(root.Left) + 1 + countNodes(root.Right)
    }
}

func leftHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + leftHeight(root.Left)
}

func rightHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + rightHeight(root.Right)
}

package main

// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/?envType=study-plan-v2&envId=leetcode-75
// 다른 사람은 head를 따로 관리하면 queue의 길이를 줄이지 않는 방법을 사용함.
func maxLevelSum(root *TreeNode) int {
    queue := []*TreeNode{root}
    minLevel := 0
    level := 1
    maxSum := -100001
    for len(queue) > 0 {
        n := len(queue)
        sum := 0
        for i := 0; i < n; i++ {
            task := queue[i]
            sum += task.Val

            if task.Left != nil {
                queue = append(queue, task.Left)
            }
            if task.Right != nil {
                queue = append(queue, task.Right)
            }
        }
        if sum > maxSum {
            minLevel = level
            maxSum = sum
        }
        queue = queue[n:]
        level++
    }
    return minLevel
}

// dfs를 사용하여 푸는 방식
func maxLevelSumDfs(root *TreeNode) int {
    mp := make(map[int]int, 0)
    recursive(root, mp, 1)
    ans := math.MinInt
    index := 0
    for key, val := range mp {
        if val > ans {
            index = key
            ans = val
        }
    }
    return index
}

func recursive(root *TreeNode, sum map[int]int, level int) {
    if root == nil {
        return
    }
    sum[level] += root.Val
    recursive(root.Left, sum, level + 1)
    recursive(root.Right, sum, level + 1)
}

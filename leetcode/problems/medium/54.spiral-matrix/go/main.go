package main

// https://leetcode.com/problems/spiral-matrix/?envType=problem-list-v2&envId=rab78cw1
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }

    n, m := len(matrix), len(matrix[0])
    answer := make([]int, n*m)
    idx := 0

    top, down := 0, n-1
    left, right := 0, m-1

    for left <= right && top <= down {
        for i := left; i <= right; i++ {
            answer[idx] = matrix[top][i]
            idx++
        }
        top++

        for i := top; i <= down; i++ {
            answer[idx] = matrix[i][right]
            idx++
        }
        right--

        if top > down || left > right {
            break
        }

        for i := right; i >= left; i-- {
            answer[idx] = matrix[down][i]
            idx++
        }
        down-- 
        
        for i := down; i >= top; i-- {
            answer[idx] = matrix[i][left]
            idx++
        }
        left++
    }

    return answer
}

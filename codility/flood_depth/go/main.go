package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
// https://app.codility.com/demo/results/training7WCQX4-S6G/
// 왼쪽 가장 높은벽과 오른쪽 가장 높은 벽을 미리 메모이제이션 해놓고 푸는 방식
func Solution(A []int) int {
	// Implement your solution here
	leftMaxHeight := make([]int, len(A))
	rightMaxHeight := make([]int, len(A))

	leftMaxHeight[0] = A[0]
	rightMaxHeight[len(A)-1] = A[len(A)-1]

	for i := 1; i < len(A); i++ {
		leftMaxHeight[i] = max(leftMaxHeight[i-1], A[i])
		rightMaxHeight[len(A)-i-1] = max(A[len(A)-i-1], rightMaxHeight[len(A)-i])
	}

	maxDepth := 0
	for i, height := range A {
		maxDepth = max(maxDepth, min(leftMaxHeight[i], rightMaxHeight[i])-height)
	}

	return maxDepth
}

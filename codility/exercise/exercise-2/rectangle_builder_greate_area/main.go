package main

/*
https://app.codility.com/c/run/training8P45UD-3FJ/
서브 어레이의 합중 가장 큰 케이스를 찾는 문제이다.
서브셋이 2개이기 때문에 처음부터 시작하는 서브셋과 뒤에서부터 시작하는 케이스를 만들었다.
음수로 내려가는 경우에는 구분을 위해서 0으로 처리했다.
서브셋이 2개이기에 쉽게 풀었지만 3개가 되는경우 어떻게 풀어야할지 많이 고민해야할거같다.
*/

func Solution(A []int) int {
	length := len(A)

	prefixSum := make([]int, length)
	for i := 1; i < length; i++ {
		prefixSum[i] = max(0, prefixSum[i-1]+A[i])
	}

	suffixSum := make([]int, length)
	for i := length - 2; i >= 0; i-- {
		suffixSum[i] = max(0, suffixSum[i+1]+A[i])
	}

	maxDoubleSliceSum := 0
	for i := 1; i < length-1; i++ {
		maxDoubleSliceSum = max(maxDoubleSliceSum, prefixSum[i-1]+suffixSum[i+1])
	}

	return maxDoubleSliceSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

// https://leetcode.com/problems/powx-n/description/
// optimization(binary Exponentiation) support from chatgtp
// Runtime 0 ms Beats 100.00% Memory 2 MB Beats 23.39%
// 절대 모든 n번만큼 모두 반복을 해버리는 경우 문제를 풀수가 없음
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	half := myPow(x, n/2)
	if n%2 == 1 {
		return half * half * x
	}
	return half * half
}

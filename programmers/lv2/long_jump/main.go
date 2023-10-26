package main

// https://school.programmers.co.kr/learn/courses/30/lessons/12914
// 뭔문제인가 하고 한참 봤다.
// 모든 경우의수를 다 계산해야하나 싶었는데 알고보니 그냥 DP문제였다.
// 1개: 1 -> 1
// 2개: 1 1,  2 -> 2
// 3개: 1 1 1, 2 1, 1 2 -> 3
// 4개: 1111, 121, 112, 211, 22 -> 5
// 5개: 11111, 122, 212, 221, 1112, 1121, 1211,  2111 -> 8
// 유사 피보나치수열 문제다.

func solution(n int) int64 {
	num := make([]int, n+1)
	num[0] = 1
	num[1] = 1
	for i := 2; i <= n; i++ {
		num[i] = (num[i-1] + num[i-2]) % 1234567
	}
	return int64(num[n])
}

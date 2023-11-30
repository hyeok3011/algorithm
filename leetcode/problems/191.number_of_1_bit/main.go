package main

/*
https://leetcode.com/problems/number-of-1-bits/?envType=daily-question&envId=2023-11-29
부호없는 이진 표현 정수를 스트링이나 이진수로 바꿔 1의 개수를 확인하면 될듯하다.

조금 신박하게 다른 방법으로 풀어보려고 비트 연산을 사용했다.
num이랑 1과 비트연산하고 한칸식 땡긴다.
Runtime 0ms Beats 100.00% Memory 1.95MB Beats 54.84%
*/
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

package main

// https://leetcode.com/problems/single-number/description/
// Runtime 17 ms Beats 67.90% Memory 6.7 MB Beats 31.12%
func singleNumber(nums []int) int {
	singleMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := singleMap[v]; ok {
			delete(singleMap, v)
		} else {
			singleMap[v] = true
		}
	}

	var answer int
	for key := range singleMap {
		answer = key
	}
	return answer
}

// 다른 사람의 풀이 이렇게 풀거라고 전혀 생각하지 못했다.
// XOR비트 연산을 활용해서 풀은듯하다
// @@@@@@@@@@@@@@@
func singleNumber2(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

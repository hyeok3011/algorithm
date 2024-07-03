package main

import "sort"

// https://leetcode.com/problems/majority-element/description/
// Runtime 15 ms Beats 85.8% Memory 6.1 MB Beats 97.15%
func majorityElement(nums []int) int {
	var majority, frequency int

	for _, v := range nums {
		if frequency == 0 {
			majority = v
			frequency = 1
		} else if majority == v {
			frequency += 1
		} else {
			frequency -= 1

		}
	}

	return majority
}

// 기본적으로 majorityElement는 n/2이 오기때문에 정렬 후 중간값을 리턴하면 답이 된다.
// 신선하게 푼듯 하다
// 정렬 과정에서 시간이 많이 걸리줄 알았는데 시간이 비슷하다... 슬라이스의 길이가 짧아서 그런가 싶기도 하고..
func anotherSolution(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

package main

import "math"

// https://leetcode.com/problems/increasing-triplet-subsequence/?envType=study-plan-v2&envId=leetcode-75
// 생각보다 어려웠던 문제
// 공간복잡도를 O(1)로 해결하는데는 결국 실패하고 다른사람의 풀이를 참고함
// @@@@@ 다시 풀어보기
func increasingTriplet(nums []int) bool {
	i := math.MaxInt32
	j := math.MaxInt32

	for k := 0; k < len(nums); k++ {
		if nums[k] <= i {
			i = nums[k]
		} else if nums[k] <= j {
			j = nums[k]
		} else {
			return true
		}
	}

	return false
}

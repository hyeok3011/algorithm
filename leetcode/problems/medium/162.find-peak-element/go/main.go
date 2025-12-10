package main

// https://leetcode.com/problems/find-peak-element/
// log N 포기
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			return i
		}
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	return -1
}

// 힌트를 보고 다시 풀어본 솔루션, @@@@@@
// 가장 높은 숫자를 찾는것이 아니라 양 사이드보다만 높으면 되는 문제이다.
// binary search의 본질은 어떠한 기준을 통해 탐색 범위 절반을 버릴 수 있는가 이다.
func findPeakElementBinarySearch(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2

		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

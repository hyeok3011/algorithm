package main

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
// Runtime 5 ms Beats 87.38% Memory 4 MB Beats 99.57%
func searchRange(nums []int, target int) []int {
	startPoint, endPoint := -1, -1

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				startPoint = mid
				break
			}
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if startPoint == -1 {
		return []int{-1, -1}
	}

	left, right = startPoint, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				endPoint = mid
				break
			}
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{startPoint, endPoint}
}

func searchRange2(nums []int, target int) []int {
	startPoint, endPoint := -1, -1

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		startPoint = left
	}

	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		endPoint = right
	}

	return []int{startPoint, endPoint}
}

package main

// https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/1852791339/?envType=study-plan-v2&envId=leetcode-75
// 가장 먼저 생각났던 방식은 당연하게도 Heap방식, 하지만 QuickSelect방식으로 풀어봄. 
// @@@@@ 다른사람의 풀이를 보다보니 정말 다양하게 푼사람들이 존재함. couting방식과 QuickSelect방식과 유사한듯한데 다른 방식, Heap을 직접 구현한 방식 등...
// 다시 풀어볼만한 문제.
func findKthLargest(nums []int, k int) int { 
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	targetIndex := n - k
	start, end := 0, n-1
	for {
		pivotIndex := partition(nums, start, end)
		if pivotIndex == targetIndex {
			break
		} else if pivotIndex < targetIndex {
			start = pivotIndex + 1
		} else {
			end = pivotIndex - 1
		}
	}

	return nums[targetIndex]
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	left, right := start+1, end

	for left <= right {
		for left <= right && nums[left] <= pivot {
			left++
		}
		for left <= right && nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right], nums[start] = nums[start], nums[right]
	return right
}



// 다른 사람의 Couting방식.. 이거는 생각하지 못했음.
func findKthLargestCouting(nums []int, k int) int {
	maxV := slices.Max(nums)
	minV := slices.Min(nums)

	couting := make([]int, maxV-minV+1)

	for _, n := range nums {
		couting[n-minV]++
	}

	for i := len(couting) - 1; i >= 0; i-- {
		k -= couting[i]

		if k <= 0 {
			return i + minV
		}
	}

	return -1
}

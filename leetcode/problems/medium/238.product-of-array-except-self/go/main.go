package main

// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75
// 생각보다 쉽지 않았던 문제...
func productExceptSelf(nums []int) []int {
	forwardProduct := make([]int, len(nums))
	forwardProduct[0] = 1
	backwardProduct := make([]int, len(nums))
	backwardProduct[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		forwardProduct[i] = forwardProduct[i-1] * nums[i-1]
		backwardProduct[len(nums)-i-1] = nums[len(nums)-i] * backwardProduct[len(nums)-i]
	}

	answer := make([]int, len(nums))
	for i := range nums {
		answer[i] = forwardProduct[i] * backwardProduct[i]
	}

	return answer
}

// 조금 더 개선 버전
func productExceptSelf2(nums []int) []int {
	answer := make([]int, len(nums))
	forwardProduct := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = forwardProduct
		forwardProduct *= nums[i]
	}

	backwardProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * backwardProduct
		backwardProduct *= nums[i]
	}

	return answer
}

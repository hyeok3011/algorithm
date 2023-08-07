package main

// https://leetcode.com/problems/maximum-product-subarray/description/
// Runtime 3 ms Beats 91.67% Memory 3.4 MB Beats 98.92%

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	totalProduct, maxProduct := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		totalProduct *= nums[i]

		maxProduct = max(totalProduct, maxProduct)

		maxProduct = max(maxProduct, nums[i])

		if totalProduct == 0 {
			totalProduct = 1 * nums[i]
		}
	}

	totalProduct = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		totalProduct *= nums[i]

		maxProduct = max(totalProduct, maxProduct)

		maxProduct = max(maxProduct, nums[i])

		if totalProduct == 0 {
			totalProduct = 1 * nums[i]
		}
	}

	return maxProduct
}

// O(2n)에서 negative 들어올때 swap하는 형식으로 하여 O(n)으로 다시 풀어봄
// 속도나 큰 차이는 없음
func maxProduct2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	bigProduct, smallProduct, maxProduct := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			bigProduct, smallProduct = smallProduct, bigProduct
		}

		bigProduct = max(nums[i], bigProduct*nums[i])
		smallProduct = min(nums[i], smallProduct*nums[i])

		maxProduct = max(maxProduct, bigProduct)
	}

	return maxProduct
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

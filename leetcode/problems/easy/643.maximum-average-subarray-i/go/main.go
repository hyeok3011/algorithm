package main

// https://leetcode.com/problems/maximum-average-subarray-i/?envType=study-plan-v2&envId=leetcode-75
func findMaxAverage(nums []int, k int) float64 {
	currentSum := sum(nums[0:k])
	maxSum := currentSum
	for i := 1; i <= len(nums)-k; i++ {
		currentSum -= nums[i-1]
		currentSum += nums[i+k-1]
		maxSum = max(maxSum, currentSum)
	}

	return float64(maxSum) / float64(k)
}
func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

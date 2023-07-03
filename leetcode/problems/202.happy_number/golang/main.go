package main

// https://leetcode.com/problems/happy-number/description/
// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 63.22%
func isHappy(n int) bool {
	sum := 0
	visited := make(map[int]bool)
	for n < 2147483647 {
		if _, ok := visited[n]; ok {
			break
		}
		visited[n] = true
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		if sum == 1 {
			return true
		}
		n = sum
		sum = 0
	}

	return false
}

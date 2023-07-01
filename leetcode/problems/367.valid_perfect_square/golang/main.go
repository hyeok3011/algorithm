package main

// https://leetcode.com/problems/valid-perfect-square/description/
// Runtime 0 ms Beats 100% Memory 1.9 MB Beats 98.40%
func isPerfectSquare(num int) bool {
	i := 1

	for i*i <= num {
		i <<= 1
	}
	i >>= 1
	for i*i <= num {
		if i*i == num {
			return true
		}
		i += 1
	}
	return false
}

// Binary search version
func isPerfectSquare2(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 2, num/2

	for left <= right {
		mid := left + (right-left)/2
		midSquared := mid * mid

		if midSquared == num {
			return true
		} else if midSquared < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

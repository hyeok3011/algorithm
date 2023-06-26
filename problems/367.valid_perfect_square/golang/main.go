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

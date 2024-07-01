package main

// https://leetcode.com/problems/count-and-say/description/
// Runtime 0 ms Beats 100% Memory 4.5 MB Beats 69.43%
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	digits := []int{1}
	for i := 0; i < n-1; i++ {
		next := []int{1, digits[0]}
		digits = digits[1:]
		for _, v := range digits {
			if v == next[len(next)-1] {
				next[len(next)-2] += 1
			} else {
				next = append(next, 1)
				next = append(next, v)
			}
		}
		digits = next
	}

	result := []byte{}
	for _, v := range digits {
		result = append(result, byte(v+'0'))
	}

	return string(result)
}

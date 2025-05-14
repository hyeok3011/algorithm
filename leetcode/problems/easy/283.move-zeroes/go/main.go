package main

// two-pointer 방식
func moveZeroes(nums []int) {
	insertPos := 0
	for i, n := range nums {
		if n != 0 {
			nums[insertPos], nums[i] = nums[i], nums[insertPos]
			insertPos++
		}
	}
}

// 무식한 방식
func moveZeroes2(nums []int) {
	for i, n := range nums {
		if n == 0 {
			j := i + 1
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j < len(nums) {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				break
			}
		}
	}
}

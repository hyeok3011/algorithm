package main

// https://leetcode.com/problems/rotate-array/description/
func rotate(nums []int, k int)  {
    k = k % len(nums)
    reverse(nums, 0, len(nums) - 1)
    reverse(nums, 0, k-1)
    reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}

func rotate2(nums []int, k int) {
    k := k % len(nums)
    if k == 0 {
        return
    }

    count := 0

    for start := 0; count < len(nums); start++ {
        current := start
        prev := nums[start]

        for {
            next := (current + k) % n
            temp := nums[next]
            nums[next] = prev
            prev = temp
            current = next
            count++

            if current == start {
                break
            }
        }
    }
}

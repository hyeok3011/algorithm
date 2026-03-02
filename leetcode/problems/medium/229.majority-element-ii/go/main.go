package main

// https://leetcode.com/problems/majority-element-ii/
// 처음 풀이, 다른사람의 majority vote 방식은 생각하지 못했었음.
func majorityElement(nums []int) []int {
    slices.Sort(nums)
    threshold := float64(len(nums) / 3)
    var step int
    if threshold < 1 {
        step = 0
    } else {
        step = int(threshold)
    }

    answer := []int{}
    i := 0
    for i < len(nums) - step {
        if nums[i] == nums[i + step] {
            answer = append(answer, nums[i])
        }
        j := i   
        for j < len(nums) - step && nums[i] == nums[j] {
            j++
        }
        i = j
    }

    return answer
}


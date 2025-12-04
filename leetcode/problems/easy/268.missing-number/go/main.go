package main

// https://leetcode.com/problems/missing-number/description/
func missingNumber(nums []int) int {
    numsSum := 0
    sum := 0
    for i, v := range nums {
        numsSum += v
        sum += i
    }
    sum += len(nums)
    return sum - numsSum
}

// 다른사람의 풀이, 가우스 선생님 방식
func missingNumber2(nums []int) int {
    s := len(nums) * (len(nums)+1) / 2

    for _, n := range nums {
        s-=n
    }

    return s
}

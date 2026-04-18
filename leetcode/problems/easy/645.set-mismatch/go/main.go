// https://leetcode.com/problems/set-mismatch/description/
// 1 ~ N까지의 숫자중에서 빠진 숫자를 찾아야한다.
// 빠진숫자가 있다는것은 중복된 숫자가 존재한다는 것이다.
// 가장 간단한 방법은 중복된 숫자를 확인하고 빠진 숫자가 확인 존재하면 중복숫자, 빠진숫자를 페어로 보내면 된다.
func findErrorNumsSet(nums []int) []int {
    set := make([]bool, len(nums) + 1)
    var duplicate int
    for _, v := range nums {
        if set[v] {
            duplicate = v
        }
        set[v] = true
    }
    var missing int
    for i := 1; i <= len(nums); i++ {
        if !set[i] {
            missing = i
            break
        }
    }
    return []int{duplicate, missing}
}

// 추가 공간을 사용하지 않고 푸는방법
func findErrorNums(nums []int) []int {
    var duplicate int
    for _, v := range nums {
        idx := abs(v) - 1
        if nums[idx] > 0 {
            nums[idx] = -nums[idx]
        } else {
            duplicate = abs(v)
        }
    }

    var missing int
    for i := range nums {
        if nums[i] > 0 {
            missing = i + 1
            break
        }
    }

    
    return []int{duplicate, missing}
}
func abs(v int) int {
    if v > 0 {
        return v
    }
    return -v
}

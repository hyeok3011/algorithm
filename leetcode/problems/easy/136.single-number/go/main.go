package main
// https://leetcode.com/problems/single-number/description/?envType=study-plan-v2&envId=leetcode-75

func singleNumber(nums []int) int {
    set := make(map[int]struct{})
    
    for _, v := range nums {
        if _, exist := set[v]; exist {
            delete(set, v)
        } else {
            set[v] = struct{}{}
        }
    }
    
    for key := range set {
        return key
    }
    
    return 0 
}

// 다른 사람의 풀이... XOR연산자 사용..
func singleNumber2(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}

package main

// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/?envType=study-plan-v2&envId=leetcode-75
// 다른사람의 최적화 방법.
// 항상 spell과 potions을 곱한뒤 비교하는 방식이 아닌 미리 sucess와 spell을 계산하여 potion이 되어야할 최솟값을 구한 뒤 비교한느 방식.
func successfulPairs(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)
    answer := make([]int, len(spells))
    potionsLen := len(potions)
    for i := range spells {
        left, right := 0, potionsLen - 1
        minPotion := (int(success) + spells[i] - 1) / spells[i]
        for left <= right {
            mid := (left + right) >> 1
            if minPotion <= potions[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
        answer[i] = potionsLen - left
    }
    return answer
}

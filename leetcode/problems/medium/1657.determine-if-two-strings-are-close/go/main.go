package main

// 다른사람의 비트연산을 사용한 사례를 보고 코드 수정... 좋은방법같음 @@@@@@@
// https://leetcode.com/problems/determine-if-two-strings-are-close/?envType=study-plan-v2&envId=leetcode-75
func closeStrings(origin string, desc string) bool {
    if len(origin) != len(desc) {
        return false    
    }
    
    var originMask, descMask uint32
    originWordsCount := [26]int{}
    descWordsCount := [26]int{}
    for i := 0; i < len(origin); i++ {
        originMask |= 1 << (origin[i] - 97)
        descMask |= 1 << (desc[i] - 97)
        originWordsCount[origin[i] - 97]++
        descWordsCount[desc[i] - 97]++
    }

    if originMask != descMask {
        return false
    }

    sort.Ints(originWordsCount[:])
    sort.Ints(descWordsCount[:])

    return originWordsCount == descWordsCount
}

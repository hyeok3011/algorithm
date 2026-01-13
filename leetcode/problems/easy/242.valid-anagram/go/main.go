package main

// https://leetcode.com/problems/valid-anagram/?envType=problem-list-v2&envId=rab78cw1
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    letters := make([]int, 26)

    for i := range len(s) {
        letters[int(s[i] - 'a')]++
    }

    for i := range len(t) {
        if letters[int(t[i] - 'a')] <= 0 {
            return false
        }
        letters[int(t[i] - 'a')]--
    }
    return true
}

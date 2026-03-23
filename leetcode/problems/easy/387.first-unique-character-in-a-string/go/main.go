package main

// https://leetcode.com/problems/first-unique-character-in-a-string/description/
func firstUniqChar(s string) int {
    var count [26]int

    for _, char := range s {
        count[char - 'a']++
    }

    for i, char := range s {
        if count[char - 'a'] == 1 {
            return i
        }
    }

    return -1
}

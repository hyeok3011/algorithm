// https://leetcode.com/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) string {
    letterLastIndex := [26]int{}
    for i, v := range s {
        letterLastIndex[v-'a'] = i
    }

    stack := []rune{}
    set := [26]bool{}
    for i, v:= range s {
        if set[v-'a'] {
            continue
        }

        j := len(stack) - 1
        for j >= 0 && stack[j] > v && letterLastIndex[stack[j]-'a'] > i {
            set[stack[j]-'a'] = false
            stack = stack[:j]
            j--
        }

        stack = append(stack, v)
        set[v-'a'] = true
    }

    return string(stack)
}

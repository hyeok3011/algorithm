package main

// https://leetcode.com/problems/group-anagrams/description/
// Runtime 24 ms Beats 79.10% Memory 7.3 MB Beats 91.32%
func groupAnagrams(words []string) [][]string {
	group := make(map[string][]string)

	for _, word := range words {
		runeLetters := make([]rune, 26)
		for _, letter := range word {
			if runeLetters[letter-97] == 0 {
				runeLetters[letter-97] = letter
			} else {
				runeLetters[letter-97] += 1
			}
		}
		if _, ok := group[string(runeLetters)]; ok {
			group[string(runeLetters)] = append(group[string(runeLetters)], word)
		} else {
			group[string(runeLetters)] = []string{word}
		}
	}

	result := [][]string{}
	for _, val := range group {
		result = append(result, val)
	}

	return result
}

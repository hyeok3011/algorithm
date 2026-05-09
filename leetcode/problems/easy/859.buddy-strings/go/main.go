package main

// https://leetcode.com/problems/buddy-strings/description/
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		letterCount := [26]bool{}
		for i := range s {
			idx := s[i] - 'a'
			if letterCount[idx] {
				return true
			}
			letterCount[idx] = true
		}
		return false
	}

	diff := make([]int, 0, 2)
	for i := range s {
		if s[i] != goal[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}

	if len(diff) != 2 {
		return false
	}

	i, j := diff[0], diff[1]
	return s[i] == goal[j] && s[j] == goal[i]
}

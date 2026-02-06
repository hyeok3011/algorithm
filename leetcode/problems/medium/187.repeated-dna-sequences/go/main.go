package main

// https://leetcode.com/problems/repeated-dna-sequences/
func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}

	sequenceSet := make(map[string]int)
	result := []string{}
	for i := range s {
		if i >= 9 {
			sequenceSet[s[i-9:i+1]]++
			if sequenceSet[s[i-9:i+1]] == 2 {
				result = append(result, s[i-9:i+1])
			}

		}
	}

	return result
}

// 다른사람의 풀이...ㄷㄷㄷ..
func findRepeatedDnaSequences2(s string) []string {
	n := len(s)
	if n < 10 {
		return nil
	}

	var enc [256]uint32
	enc['A'] = 0
	enc['C'] = 1
	enc['G'] = 2
	enc['T'] = 3

	const mask uint32 = (1 << 20) - 1

	counts := make(map[uint32]uint8, n-9)

	var code uint32
	for i := 0; i < 10; i++ {
		code = (code << 2) | enc[s[i]]
	}
	counts[code] = 1

	res := make([]string, 0, 16)

	for i := 10; i < n; i++ {
		code = ((code << 2) | enc[s[i]]) & mask

		c := counts[code]
		if c == 1 {
			res = append(res, s[i-9:i+1])
			counts[code] = 2
		} else if c == 0 {
			counts[code] = 1
		}
	}

	return res
}

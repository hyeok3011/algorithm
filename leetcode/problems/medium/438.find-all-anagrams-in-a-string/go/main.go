// https://leetcode.com/problems/find-all-anagrams-in-a-string/?envType=problem-list-v2&envId=rab78cw1
// @@@@@@ 방법은 알았는데 이상하게 구현하는데 오래 걸린 문제
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	var pCount [26]int
	var sCount [26]int

	for i := range p {
		pCount[p[i]-'a']++
	}

	answer := []int{}
	matchCount := 0

	for i := 0; i < len(p); i++ {
		index := s[i] - 'a'
		sCount[index]++
		if sCount[index] <= pCount[index] {
			matchCount++
		}
	}

	if matchCount == len(p) {
		answer = append(answer, 0)
	}

	for i := 0; i < len(s)-len(p); i++ {
		removeIndex := s[i] - 'a'
		sCount[removeIndex]--
		if sCount[removeIndex] < pCount[removeIndex] {
			matchCount--
		}

		addIndex := s[i+len(p)] - 'a'
		sCount[addIndex]++
		if sCount[addIndex] <= pCount[addIndex] {
			matchCount++
		}

		if matchCount == len(p) {
			answer = append(answer, i+1)
		}
	}

	return answer
}

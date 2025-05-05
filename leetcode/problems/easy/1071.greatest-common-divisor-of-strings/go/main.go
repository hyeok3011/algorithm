package main

// https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75
func gcdOfStrings(str1 string, str2 string) string {
	length1, length2 := len(str1), len(str2)
	maxPrefixLength := length1
	if length2 < length1 {
		maxPrefixLength = length2
	}

	for currentPrefixLength := maxPrefixLength; currentPrefixLength >= 1; currentPrefixLength-- {
		if length1%currentPrefixLength == 0 && length2%currentPrefixLength == 0 {
			candidatePrefix := str1[:currentPrefixLength]

			if checkComposition(str1, candidatePrefix) && checkComposition(str2, candidatePrefix) {
				return candidatePrefix
			}
		}
	}

	return ""
}
func checkComposition(s string, prefix string) bool {
	stringLength := len(s)
	prefixLength := len(prefix)

	if prefixLength == 0 {
		return stringLength == 0
	}

	if stringLength%prefixLength != 0 {
		return false
	}

	repeatCount := stringLength / prefixLength
	for i := 0; i < repeatCount; i++ {
		substringToCheck := s[i*prefixLength : (i+1)*prefixLength]
		if substringToCheck != prefix {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/greatest-common-divisor-of-strings/solutions/6601478/euclid-s-algorithm-visualized-step-by-step-gcd-mastery-beginner-friendly-guide/?envType=study-plan-v2&envId=leetcode-75
// 다른사람의 유클리드 방법 ㄷㄷ... 다시 문제 확인해보니 충분히 유추 가능했음...
func gcdOfStrings2(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	a, b := len(str1), len(str2)
	for b > 0 {
		a, b = b, a%b
	}

	return str1[:a]
}

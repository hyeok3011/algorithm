package main

func reverseVowels(s string) string {
	vowelsSet := make(map[string]bool)
	for _, v := range []string{"A", "a", "E", "e", "I", "i", "O", "o", "U", "u"} {
		vowelsSet[v] = true
	}

	runes := []rune(s)
	start := 0
	end := len(runes) - 1
	for start < end {
		for start < end && !vowelsSet[string(runes[start])] {
			start++
		}

		for start < end && !vowelsSet[string(runes[end])] {
			end--
		}

		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}

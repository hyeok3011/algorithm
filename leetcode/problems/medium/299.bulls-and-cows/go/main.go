package main

import "fmt"

// Bull (A, 황소): 숫자도 맞고 자리도 맞음
// Cow (B, 소): 숫자는 있는데 자리가 틀림
// https://leetcode.com/problems/bulls-and-cows/
func getHint(secret string, guess string) string {
	bullCount := 0
	secretCount := [10]int{}
	guessCount := [10]int{}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bullCount++
		} else {
			secretCount[secret[i]-'0']++
			guessCount[guess[i]-'0']++
		}
	}

	cowCount := 0
	for i := 0; i < len(guessCount); i++ {
		if guessCount[i] > 0 {
			cowCount += min(secretCount[i], guessCount[i])
		}
	}

	return fmt.Sprintf("%dA%dB", bullCount, cowCount)
}

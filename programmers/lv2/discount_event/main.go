package main

// https://school.programmers.co.kr/learn/courses/30/lessons/131127
func solution(want []string, number []int, discount []string) int {
	wantItemCount := make(map[string]int, 10)

	for i, item := range want {
		wantItemCount[item] = number[i]
	}

	for i := 0; i < 10; i++ {
		if _, exist := wantItemCount[discount[i]]; exist {
			wantItemCount[discount[i]]--
		}
	}

	days := 0
	if isCompleted(wantItemCount) {
		days += 1
	}
	for i := 10; i < len(discount); i++ {
		if _, exist := wantItemCount[discount[i-10]]; exist {
			wantItemCount[discount[i-10]]++
		}

		if _, exist := wantItemCount[discount[i]]; exist {
			wantItemCount[discount[i]]--
		}
		if isCompleted(wantItemCount) {
			days += 1
		}
	}

	return days
}

func isCompleted(wantItemCount map[string]int) bool {
	for _, val := range wantItemCount {
		if val != 0 {
			return false
		}
	}

	return true
}

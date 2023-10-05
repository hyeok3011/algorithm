package main

// https://school.programmers.co.kr/learn/courses/30/lessons/77484
// winningNumber을 set에 넣어둬 O(1)로 확인할수 있도록 준비하고 맞춘 볼카운트와 0개수를 세면 된다
// 최고 순위는 맞춘 볼 카운트와 zerocount를 합치고 최저순위는 볼카운트만 생각하면 되는문제
func solution(lottos []int, win_nums []int) []int {
	winningNumber := make(map[int]bool)

	for i := range win_nums {
		winningNumber[win_nums[i]] = true
	}

	zeroCount := 0
	matchedCount := 0
	for _, v := range lottos {
		if v == 0 {
			zeroCount++
		} else if winningNumber[v] {
			matchedCount++
		}
	}

	highestRank := 7 - matchedCount - zeroCount
	if highestRank > 6 {
		highestRank = 6
	}
	lowestRank := 7 - matchedCount
	if lowestRank > 6 {
		lowestRank = 6
	}

	return []int{highestRank, lowestRank}
}

// rank를 map이나 switch로 풀면 깔끔할거같음
func solution2(lottos []int, win_nums []int) []int {
	winningNumber := make(map[int]bool)

	for i := range win_nums {
		winningNumber[win_nums[i]] = true
	}

	zeroCount := 0
	matchedCount := 0
	for _, v := range lottos {
		if v == 0 {
			zeroCount++
		} else if winningNumber[v] {
			matchedCount++
		}
	}

	highestRank := getRank(matchedCount + zeroCount)
	lowestRank := getRank(matchedCount)

	return []int{highestRank, lowestRank}
}

func getRank(score int) int {
	switch score {
	case 6:
		return 1
	case 5:
		return 2
	case 4:
		return 3
	case 3:
		return 4
	case 2:
		return 5
	default:
		return 6
	}
}

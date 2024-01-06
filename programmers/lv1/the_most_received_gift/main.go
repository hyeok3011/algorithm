package main

import "strings"

/*
https://school.programmers.co.kr/learn/courses/30/lessons/258712
*/

func solution(friends []string, gifts []string) int {
	friendGiftIndex := make(map[string]int, len(friends))
	friendSendGift := make(map[string]map[string]int, len(friends))

	for _, gift := range gifts {
		splited := strings.Split(gift, " ")
		from, to := splited[0], splited[1]
		friendGiftIndex[from]++
		friendGiftIndex[to]--

		if _, exist := friendSendGift[from]; !exist {
			friendSendGift[from] = make(map[string]int)
		}

		friendSendGift[from][to]++
	}

	friendReceiveCount := make([]int, len(friends))
	for i := 0; i < len(friends); i++ {
		from := friends[i]
		for j := i + 1; j < len(friends); j++ {
			to := friends[j]
			if friendSendGift[from][to] > friendSendGift[to][from] {
				friendReceiveCount[i]++
			} else if friendSendGift[from][to] < friendSendGift[to][from] {
				friendReceiveCount[j]++
			} else if friendGiftIndex[from] > friendGiftIndex[to] {
				friendReceiveCount[i]++
			} else if friendGiftIndex[from] < friendGiftIndex[to] {
				friendReceiveCount[j]++
			}
		}
	}

	maxCount := 0
	for _, v := range friendReceiveCount {
		if maxCount < v {
			maxCount = v
		}
	}
	return maxCount
}

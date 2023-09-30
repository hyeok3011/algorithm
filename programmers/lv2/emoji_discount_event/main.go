package main

import "log"

// https://school.programmers.co.kr/learn/courses/30/lessons/150368
// 문제 이해하는게 오래걸렸음... 카카오 문제는 이해하는데 너무 힘듬...
func solution(users [][]int, emoticons []int) []int {
	plusCount, sellPrice := dfs(users, emoticons, make([]int, len(emoticons)), 0)
	return []int{plusCount, sellPrice}
}

func dfs(users [][]int, emoticons, emoticonDiscount []int, emoticonIndex int) (int, int) {
	maxPlusCount, maxProfit := 0, 0

	if emoticonIndex == len(emoticons) {
		log.Println(emoticonDiscount)
		for i := range users {
			userDiscountRate, userMaximumPrice := users[i][0], users[i][1]
			profit := 0
			for j, emoticonPrice := range emoticons {
				if emoticonDiscount[j] >= userDiscountRate {
					profit += int(float32(emoticonPrice) * (100 - float32(emoticonDiscount[j])) / 100)
				}
			}

			if userMaximumPrice <= profit {
				maxPlusCount++
			} else {
				maxProfit += profit
			}
		}
		return maxPlusCount, maxProfit
	}

	for _, discountRate := range []int{10, 20, 30, 40} {
		emoticonDiscount[emoticonIndex] = discountRate
		plustCount, profit := dfs(users, emoticons, emoticonDiscount, emoticonIndex+1)
		if plustCount > maxPlusCount {
			maxPlusCount = plustCount
			maxProfit = profit
		} else if plustCount == maxPlusCount && maxProfit < profit {
			maxProfit = profit
		}
	}

	return maxPlusCount, maxProfit
}

// chan version
// 다른사람이 알고리즘 문제에도 channel을 쓰길래 대충 써봄...
func solution2(users [][]int, emoticons []int) []int {
	maxPlusCount, maxProfit := 0, 0
	discountChan := make(chan []int)
	go func() {
		generateDiscountCombinationsDFS(make([]int, len(emoticons)), 0, discountChan)
		close(discountChan)
	}()
	for emoticonDiscount := range discountChan {
		plusCount, profit := 0, 0
		for i := range users {
			userDiscountRate, userMaximumPrice := users[i][0], users[i][1]
			totalPrice := 0
			for j, emoticonPrice := range emoticons {
				if emoticonDiscount[j] >= userDiscountRate {
					totalPrice += int(float32(emoticonPrice) * (100 - float32(emoticonDiscount[j])) / 100)
				}
			}

			if userMaximumPrice <= totalPrice {
				plusCount++
			} else {
				profit += totalPrice
			}
		}
		if plusCount > maxPlusCount {
			maxPlusCount = plusCount
			maxProfit = profit
		} else if plusCount == maxPlusCount && maxProfit < profit {
			maxProfit = profit
		}
	}

	return []int{maxPlusCount, maxProfit}
}

func generateDiscountCombinationsDFS(emoticonDiscount []int, emoticonIndex int, discountChan chan []int) {
	if emoticonIndex == len(emoticonDiscount) {
		discountRateCopy := make([]int, len(emoticonDiscount))
		copy(discountRateCopy, emoticonDiscount)
		discountChan <- discountRateCopy
	} else {
		for _, discountRate := range []int{10, 20, 30, 40} {
			emoticonDiscount[emoticonIndex] = discountRate
			generateDiscountCombinationsDFS(emoticonDiscount, emoticonIndex+1, discountChan)
		}
	}
}

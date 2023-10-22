package main

import (
	"sort"
	"time"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/155651
// 먼저 입장 시간으로 오름차순으로 정렬해놓고

func solution(bookTimes [][]string) int {
	orders := make([][]time.Time, len(bookTimes))
	for i := range bookTimes {
		startTime := parseTime(bookTimes[i][0])
		endTime := parseTime(bookTimes[i][1])
		orders[i] = []time.Time{startTime, endTime}
	}

	sort.Slice(orders, func(i, j int) bool {
		return orders[i][0].Before(orders[j][0])
	})

	endTimeOfRooms := []time.Time{}
	for i := range orders {
		isNotRentalRoom := true
		for j := range endTimeOfRooms {
			if endTimeOfRooms[j].Before(orders[i][0]) || endTimeOfRooms[j].Equal(orders[i][0]) {
				endTimeOfRooms[j] = orders[i][1].Add(time.Minute * 10)
				isNotRentalRoom = false
				break
			}
		}

		if isNotRentalRoom {
			endTimeOfRooms = append(endTimeOfRooms, orders[i][1].Add(time.Minute*10))
		}
	}

	return len(endTimeOfRooms)
}

func parseTime(strTime string) time.Time {
	t, _ := time.Parse("15:04", strTime)
	return t
}

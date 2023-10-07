package main

import (
	"sort"
	"strings"
	"time"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/92341
// 입차 출차 분을 계산하여 간단한 연산문제 이다.
// 나는 time.Parse를 사용했지만 단순하게 그냥 시간과 분을 int로 보고 total 분만 계산하는게 더 편했을거 같다.
// 올림 계산을 math라이브러리를 사용하지않고 단위를 하나더 더한뒤 -1하는식으로 풀었다.
func solution(fees []int, records []string) []int {
	defaultParkingMinut := fees[0]
	defaultParkingFee := fees[1]
	minutUnit := fees[2]
	feePerMinutUnit := fees[3]

	carInOutTime := make(map[string][]time.Time)
	for _, v := range records {
		splited := strings.Split(v, " ")
		carNumber := splited[1]
		if _, exist := carInOutTime[carNumber]; exist {
			carInOutTime[carNumber] = append(carInOutTime[carNumber], parseTime(splited[0]))
		} else {
			carInOutTime[carNumber] = []time.Time{parseTime(splited[0])}
		}
	}

	keys := make([]string, 0, len(carInOutTime))
	for k := range carInOutTime {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	answer := make([]int, len(carInOutTime))
	for i, carNumber := range keys {
		totalParkingFee := defaultParkingFee
		totalMinute := 0
		if len(carInOutTime[carNumber])%2 == 1 {
			carInOutTime[carNumber] = append(carInOutTime[carNumber], parseTime("23:59"))
		}
		for j := 0; j < len(carInOutTime[carNumber]); j += 2 {
			totalMinute += int(carInOutTime[carNumber][j+1].Sub(carInOutTime[carNumber][j]).Minutes())
		}
		if totalMinute > defaultParkingMinut {
			totalParkingFee += ((totalMinute - defaultParkingMinut + minutUnit - 1) / minutUnit) * feePerMinutUnit
		}
		answer[i] = totalParkingFee
	}

	return answer
}

func parseTime(strTime string) time.Time {
	t, _ := time.Parse("15:04", strTime)
	return t
}

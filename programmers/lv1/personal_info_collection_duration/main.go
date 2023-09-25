package main

import (
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT = "2006.01.02"

// https://school.programmers.co.kr/learn/courses/30/lessons/150370
func solution(strToday string, terms []string, privacies []string) []int {
	policyTerms := make(map[string]int)

	for i := range terms {
		splited := strings.Split(terms[i], " ")
		policyType := splited[0]
		policyTerm, _ := strconv.Atoi(splited[1])
		policyTerms[policyType] = policyTerm
	}
	today, _ := time.Parse(DATE_FORMAT, strToday)

	answer := []int{}
	for i := range privacies {
		splited := strings.Split(privacies[i], " ")
		collectedDate, _ := time.Parse(DATE_FORMAT, splited[0])

		expireDate := collectedDate.AddDate(0, policyTerms[splited[1]], 0)
		if expireDate.Before(today) || expireDate.Equal(today) {
			answer = append(answer, i+1)
		}
	}

	return answer
}

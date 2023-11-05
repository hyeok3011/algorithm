package main

import (
	"sort"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/42746
// golang의 string은 유니코드를 사용한다. https://symbl.cc/en/unicode/table/
// i + j 로 비교한느 이뉴는 단순 i > j 형식으로 비교하게되면 3 < 30같은 상황이 나오게 된다.
func solution(numbers []int) string {

	strNums := make([]string, len(numbers))
	for i, num := range numbers {
		strNums[i] = strconv.Itoa(num)
	}

	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	if strNums[0] == "0" {
		return "0"
	}

	var maximumBulder strings.Builder

	for _, v := range strNums {
		maximumBulder.WriteString(v)
	}

	return maximumBulder.String()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

// https://www.hackerrank.com/challenges/non-divisible-subset/problem?isFullScreen=true
// 요소 2개의 수를 더하더라도 k로 나눠지지 않는 가장 긴 부분집합을 구하라는 뜻이다.
// n은 10**5까지 되기때문에 단순 조합으로는 정답을 구하기 힘들듯 하다.
// 예제처럼 나머지의 합이 k가 되지 않는 조합으로 하면 될거같다
// s의 모든 요소들의 나머지 개수를 구하고
// 나머지 두개의 합이 k가 되는 케이스중 더 높은 점수를 선택해야한다.
// 풀다보니 예외가 더 있다.
// 0은 1개까지만 가능 왜냐하면 0이 2개가되면 문제 발생한다.
// 그리고 k가 짝수인경우  1개까지만 가능 왜냐하면 나머지가 k/2 가 2개 이상이 되는경우 문제가 발생한다.

// 최근에 푼 문제중 시간좀 좀 걸리편 같다.
func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	remainderCount := make(map[int32]int32, len(s))
	for i := range s {
		remainderCount[s[i]%k] += 1
	}

	var subsetCount int32
	if remainderCount[0] > 0 {
		subsetCount += 1
	}

	if k%2 == 0 && remainderCount[k/2] > 0 {
		subsetCount += 1
		remainderCount[k/2] = 0
	}

	//  1 2 3 4 5 6
	for i := int32(1); i <= k/2; i++ {
		subsetCount += max(remainderCount[i], remainderCount[k-i])
	}
	return subsetCount
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

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
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

/*
https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem?isFullScreen=true
문제 자체는 어렵지 않다. 다만 길이가 ranked와 player의 길이가 10 ** 5라는 점이다.
단순하게 player순회하며 ranked를 확인하면 O(mn)이다.
그래서 고민한것이
1. ranked를 중복 제거 후 player를 순회할때마다 ranked를 이진탐색으로 확인하는 방법 = O(n + n Log m)
2. 중복 제거 후 player를 순회하는동안 중복제거된 ranked를 한번만 순회하는 방법 = O(2n + m)
흠... 먼저 2번 방법으로 풀어본다.

일단 풀리기는 했는데 테이스를 만들어서 두 방법의 속도를 비교해보고 싶지만 너무 귀찮다.
*/

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	lastUniqueRankIndex := int32(0)
	uniqueRankedScores := []int32{2147483647}
	for i := range ranked {
		if uniqueRankedScores[lastUniqueRankIndex] > ranked[i] {
			lastUniqueRankIndex++
			uniqueRankedScores = append(uniqueRankedScores, ranked[i])
		}
	}

	playerRanks := make([]int32, len(player))
	for i, score := range player {
		for lastUniqueRankIndex >= 0 && score >= uniqueRankedScores[lastUniqueRankIndex] {
			lastUniqueRankIndex--
		}

		playerRanks[i] = lastUniqueRankIndex + 1
	}
	return playerRanks
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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

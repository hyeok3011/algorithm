package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

// https://www.hackerrank.com/challenges/magic-square-forming/problem?isFullScreen=true
// 마방진 문제는 처음보는것 같다.
// 쉽게 정답을 만들수 있는 패턴이 있을까 계속 찾아봤지만 결국 찾지 못했다.
// 결국 마방진이 될수있는 모든 경우를 찾아 비교하는 방법뿐이 없을듯 하다.
// 모든 경우의수를 확인했으나 3 * 3 고정이니 그냥 처음부터 가능한 마방진을 사용하는게 더 효율적이였음...

func formingMagicSquare(s [][]int32) int32 {
	allMagicSquares := [8][3][3]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}

	minimum := int32(math.MaxInt32)
	for _, magicsquare := range allMagicSquares {
		var diff int32
		for rowIndex := range magicsquare {
			for colIndex := range magicsquare[rowIndex] {
				diff += abs(s[rowIndex][colIndex] - magicsquare[rowIndex][colIndex])
			}
		}
		minimum = min(minimum, diff)
	}

	return minimum
}

func min(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

func abs(v int32) int32 {
	if v > 0 {
		return v
	}
	return -v
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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

// 4 9 2 15
// 3 5 7 15
// 8 1 5 14
// 151514 14

// 5 3 4 12
// 1 5 8 14
// 6 4 2 12
//121214 12

// 4 8 2 14
// 4 5 7 16
// 6 1 6 13
//141414
//

// 4 9 2 15
// 3 5 7 15
// 8 1 6 13
//141515

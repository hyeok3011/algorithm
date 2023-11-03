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
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */
// https://www.hackerrank.com/challenges/array-left-rotation/problem?isFullScreen=true
func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	// n - i - d
	// 4 - 4
	// n - 0 - 4
	rotated := make([]int32, len(arr))
	for i, v := range arr {
		newIndex := int32(i) - d
		if newIndex < 0 {
			newIndex = int32(len(arr)) + newIndex
		}
		rotated[newIndex] = v
	}

	return rotated
}
func reverse(arr []int32, start, end int32) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// inplace방법
func rotateLeft2(d int32, arr []int32) []int32 {
	n := int32(len(arr))
	if d == 0 || d == n {
		return arr
	}

	d = d % n // 무의미하게 큰 경우
	// 12345
	reverse(arr, 0, d-1) // 4 3 2 1 5 // 2 1 3 4 5

	reverse(arr, d, n-1) // //  2 1 5 4 3

	reverse(arr, 0, n-1) // 5 1 2 3 4 // 3 4 5 1 2

	return arr
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

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := rotateLeft(d, arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
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

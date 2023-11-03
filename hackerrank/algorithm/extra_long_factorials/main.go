package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */
//  https://hackerrank.com/challenges/extra-long-factorials/problem?isFullScreen=true
// 100!을 확인해보니 자리수가 그냥 말이되는듯 하다 159였나
// 팩토리얼 연산을 줄일수 있는 방법을 찾아봤으나 없는듯 하다... 어째되었든 모든 연산은 수행해야 한다.
// 그래서 생각한것이 숫자들의 자릿수를 slice에 하나씩 담아 연산하려고 했으나
// 자릿수가 늘어날수록 비효율적일듯 싶다....
// golang에 bigInt가있고 심지어 MulRange가 있다.....
// MulRange를 확인해보니 재귀로 분할정복 알고리즘을 사용하고 있는것으로 확인됐다.

func extraLongFactorials(n int32) {
	result := big.Int{}
	// Write your code here
	result.MulRange(1, int64(n))

	fmt.Println(result.String())
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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

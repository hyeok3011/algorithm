package main

/*
https://leetcode.com/problems/gray-code/description/
사실 문제만 봤을때는 감이 잡히지 않아 gray code가 뭔지 위키문서를 확인해봤다.
Gray Code : https://en.wikipedia.org/wiki/Gray_code
머쓱하게도 위 문서에 i의 gray code를 구하는 방법이 바로 나와있어 나머지는 문제 요구사항에 맞춰 풀었다.
*/

func grayCode(n int) []int {
	length := 1
	for i := 0; i < n; i++ {
		length <<= 1
	}

	answer := make([]int, length)

	for i := 0; i < length; i++ {
		answer[i] = i ^ (i >> 1)
	}

	return answer
}

package main

import "fmt"

// https://leetcode.com/problems/multiply-strings/description/
// Runtime 3 ms Beats 73.33% Memory 3.2 MB Beats 50.48%
// Code refactoring and optimization in progress.

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	firstFactorArray := convertStrIntegerToIntgerArray(num1)
	secondFactorArray := convertStrIntegerToIntgerArray(num2)

	multipleArray := make([]int, len(num1)+len(num2))
	length := len(multipleArray) - 1
	for secondFactorIndex, secondFactorVal := range secondFactorArray {
		for firstFactorIndex, firstFactorVal := range firstFactorArray {
			currentIndex := length - secondFactorIndex - firstFactorIndex
			multiple := (secondFactorVal * firstFactorVal)
			multipleArray[currentIndex] += multiple
			if multipleArray[currentIndex] > 9 {
				multipleArray[currentIndex-1] += multipleArray[currentIndex] / 10
				multipleArray[currentIndex] = multipleArray[currentIndex] % 10
			}
		}
	}

	result := ""
	if multipleArray[0] == 0 {
		multipleArray = multipleArray[1:]
	}
	for _, v := range multipleArray {
		result += fmt.Sprint(v)
	}
	return result
}

func convertStrIntegerToIntgerArray(strInt string) []int {
	integerArray := make([]int, len(strInt))
	length := len(strInt) - 1
	for index, v := range strInt {
		integerArray[length-index] = int(v - 48)
	}
	return integerArray
}

// 123
// 456

// 492
//  615
//   738

// 18 12 6
// 15 10 5
// 12 8 4

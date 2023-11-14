package main

// func solution(outTemperature, lowTemperature, highTemperature, diffPower, samePower int, onboard []int) int {
// 	outTemperature, lowTemperature, highTemperature = outTemperature+10, lowTemperature+10, highTemperature+10

// 	var customerCount int
// 	for i := range onboard {
// 		if onboard[i] == 1 {
// 			customerCount++
// 		}
// 	}

// 	var totalPower, tempDiff int

// 	startIndex := 0
// 	for i := range onboard {
// 		if onboard[i] == 1 {
// 			startIndex = i
// 			totalPower += max(lowTemperature-outTemperature, outTemperature-highTemperature) * diffPower
// 			break
// 		}
// 	}
// 	for i :=
// 		startIndex; i < len(onboard); i++ {
// 		if onboard[i] == 1 {
// 			customerCount--
// 		}

// 		if customerCount > 0 {
// 			if tempDiff > 0 {
// 				totalPower += diffPower
// 				tempDiff -= 2
// 			} else if samePower*2 < diffPower && tempDiff == 0 || (customerCount == 1 && onboard[i+1] == 1 && diffPower > samePower && tempDiff == 0) {
// 				totalPower += samePower
// 				tempDiff--
// 			} else if tempDiff >= 0 && samePower*2 >= diffPower {
// 				totalPower += diffPower
// 				tempDiff -= 2
// 			}
// 		}
// 		tempDiff++
// 	}

// 	return totalPower
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }

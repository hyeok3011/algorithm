package main

func addStrings(num1 string, num2 string) string {
	maxLen := max(len(num1), len(num2))

	num1Index := len(num1) - 1
	num2Index := len(num2) - 1
	result := make([]byte, maxLen+1)
	index := len(result) - 1
	var overflow byte
	for num1Index >= 0 || num2Index >= 0 || overflow != 0 {
		sum := overflow
		if num1Index >= 0 {
			sum += num1[num1Index] - '0'
			num1Index--
		}
		if num2Index >= 0 {
			sum += num2[num2Index] - '0'
			num2Index--
		}

		result[index] = sum%10 + '0'
		overflow = sum / 10
		index--
	}

	if result[0] == 0 {
		return string(result[1:])
	}

	return string(result)
}

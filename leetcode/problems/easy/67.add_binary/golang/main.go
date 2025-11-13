package main

func addBinary(a string, b string) string {
	aReversedbyteArray := stringToReverseByteArray(a)
	bReversedbyteArray := stringToReverseByteArray(b)
	if len(aReversedbyteArray) < len(bReversedbyteArray) {
		bReversedbyteArray, aReversedbyteArray = aReversedbyteArray, bReversedbyteArray
	}

	bLen := len(bReversedbyteArray) - 1
	carry := byte(0)
	for index, _ := range aReversedbyteArray {
		aReversedbyteArray[index] += carry
		carry = 0

		if bLen >= index {
			aReversedbyteArray[index] += bReversedbyteArray[index]
		}

		if aReversedbyteArray[index] > 1 {
			carry = aReversedbyteArray[index] / 2
			aReversedbyteArray[index] = aReversedbyteArray[index] % 2
		}
	}

	result := ""
	if carry == 1 {
		result = "1"
	}
	for i := len(aReversedbyteArray) - 1; i >= 0; i-- {
		result += string(aReversedbyteArray[i] + 48)
	}

	return result
}

func stringToReverseByteArray(src string) []byte {
	byteArray := make([]byte, len(src))
	srcLen := len(src) - 1
	for index, v := range src {
		byteArray[srcLen-index] = byte(v - 48)
	}

	return byteArray
}

func addBinary2(a string, b string) string {
    aIndex, bIndex := len(a) - 1, len(b) - 1
    overflow := false
    sumBuffer := make([]byte, max(len(a), len(b)))
    index := len(sumBuffer) - 1
    for aIndex >= 0 || bIndex >= 0 {
        sum := 0
        if aIndex >= 0{
            sum += int(a[aIndex] - '0')
            aIndex--
        }
        if bIndex >= 0{
            sum += int(b[bIndex] - '0')
            bIndex--
        }
        if overflow {
            sum += 1
            overflow = false
        }
        overflow = sum > 1
        sumBuffer[index] = byte(sum % 2 + '0')
        index--
    }
    answer := string(sumBuffer)
    if overflow {
        answer = "1" + answer
    }
    return answer
}

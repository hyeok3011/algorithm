// https://leetcode.com/problems/decode-string/description/
func decodeString(s string) string {
	var recursion func(start int) ([]byte, int)
	recursion = func(start int) ([]byte, int) {
		result := []byte{}
		repeat := 0
		var i int
		for i = start; i < len(s); i++ {
			if '0' <= s[i] && s[i] <= '9' {
				if repeat != 0 {
					repeat *= 10
				}
				repeat += int(s[i] - '0')
			} else if s[i] == '[' {
				subChar, newIndex := recursion(i + 1)
				for j := 0; j < repeat; j++ {
					result = append(result, subChar...)
				}
				i = newIndex
				repeat = 0
			} else if s[i] == ']' {
				break
			} else {
				result = append(result, s[i])
			}
		}
		start = i
		return result, start
	}

	answer, _ := recursion(0)

	return string(answer)
}


type stackFrame struct {
	repeat int
	buf    []byte
}

func decodeStringStack(str string) string {
	stack := []stackFrame{stackFrame{}}
	for _, v := range str {
		if '0' <= v && v <= '9' {
			if stack[len(stack)-1].repeat != 0 {
				stack[len(stack)-1].repeat *= 10
			}
			stack[len(stack)-1].repeat += int(v - '0')
		} else if v == '[' {
			stack = append(stack, stackFrame{})
		} else if v == ']' {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for j := 0; j < stack[len(stack)-1].repeat; j++ {
				stack[len(stack)-1].buf = append(stack[len(stack)-1].buf, temp.buf...)
			}
			stack[len(stack)-1].repeat = 0
		} else {
			stack[len(stack)-1].buf = append(stack[len(stack)-1].buf, byte(v))
		}
	}

	return string(stack[0].buf)
}

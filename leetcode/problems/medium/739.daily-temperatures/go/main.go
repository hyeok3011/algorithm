// https://leetcode.com/problems/daily-temperatures/description/
// @@@@ 이거 다시봐야하는 문제. cpu cache관련 내용
func dailyTemperatures(temperatures []int) []int {
    answer := make([]int, len(temperatures))

    stack := make([]int, 0, len(temperatures))
    for i, temp := range temperatures {
        for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < temp {
            pendingIndex := stack[len(stack) - 1]
            answer[pendingIndex] = i - pendingIndex
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
    }
    return answer
}

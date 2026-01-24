package main

// https://leetcode.com/problems/largest-number/
func largestNumber(nums []int) string {
    numStrs := make([]string, len(nums))
    for i := range nums {
        numStrs[i] = strconv.Itoa(nums[i])
    }
    sort.Slice(numStrs, func(i, j int) bool {
        return numStrs[i] + numStrs[j] > numStrs[j] + numStrs[i]
    })

    if numStrs[0] == "0" {
        return "0"
    }

    var sb strings.Builder
    for i := range numStrs {
        sb.WriteString(numStrs[i])
    }
    return sb.String()
}

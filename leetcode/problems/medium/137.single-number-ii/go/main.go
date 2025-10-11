// https://leetcode.com/problems/single-number-ii/description/
// @@@@@@ 접근할때 힌트 얻고 접근했고 오랜만에 bit 푸는데도 헷갈렸음.....
func singleNumber(nums []int) int {
	numsBits := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < len(numsBits); i++ {
			bit := (num >> i) & 1
			numsBits[i] += bit
		}
	}

	var answer int
	for i := 0; i < len(numsBits); i++ {
		if numsBits[i]%3 == 1 {
			answer += 1 << i
		}

	}
	return int(int32(answer))
}

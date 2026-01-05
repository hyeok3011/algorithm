// https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/?envType=study-plan-v2&envId=leetcode-75
func minFlips(a int, b int, c int) int {
	if a+b == c {
		return 0
	}

	swap := 0
	for i := 0; i <= 31; i++ {
		aBit := (a >> i) & 1
		bBit := (b >> i) & 1
		cBit := (c >> i) & 1
		if cBit == 0 && (aBit != 0 || bBit != 0) {
			swap += aBit + bBit
		} else if cBit == 1 != (aBit == 1 || bBit == 1) {
			swap += 1
		}
	}

	return swap
}

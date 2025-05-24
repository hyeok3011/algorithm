package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

const (
	Higher = -1
	Equal  = 0
	Lower  = 1
)

func guessNumber(n int) int {
	left := 1
	right := n
	var mid int
	for {
		mid = (left + right) / 2
		guessResult := guess(mid)
		if guessResult == Equal {
			break
		} else if guessResult == Higher {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return mid
}

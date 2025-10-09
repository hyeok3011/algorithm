// https://leetcode.com/problems/first-bad-version/description/

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    for isBadVersion(n - 1) {
        n--
    }

    return n
}

func firstBadVersionTwoPointer(n int) int {
	left := 1
	right := n
	startBadVersion := n
	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			startBadVersion = mid
			right = mid
		} else {
			left = mid + 1
		}
	}
	return startBadVersion
}

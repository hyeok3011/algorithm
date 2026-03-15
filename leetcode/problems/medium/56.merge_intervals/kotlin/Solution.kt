// https://leetcode.com/problems/merge-intervals/
class Solution {
    fun merge(intervals: Array<IntArray>): Array<IntArray> {
        intervals.sortBy{it[0]}
        
        val answer = mutableListOf<IntArray>()
        var prev = intervals[0]
        for (i in 1 until intervals.size) {
            val curr = intervals[i]
            if (curr[0] <= prev[1]) {
                prev[1] = maxOf(prev[1], curr[1])
            } else {
                answer.add(prev)
                prev = curr
            }
        }

        answer.add(prev)
        return answer.toTypedArray()
    }
}
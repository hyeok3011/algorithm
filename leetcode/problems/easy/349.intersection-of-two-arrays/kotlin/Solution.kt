// https://leetcode.com/problems/intersection-of-two-arrays/
class Solution {
    fun intersection(nums1: IntArray, nums2: IntArray): IntArray {
        val seen = mutableSetOf<Int>()
        for (num in nums1) {
            seen.add(num)
        }

        val result = mutableSetOf<Int>()
        for (num in nums2) {
            if (num in seen) {
                result.add(num)
            }
        }

        return result.toIntArray()
    }
    
    // 다른사람의 풀이, 최소 사이즈 잡아놓고 잘라서 반환
    fun intersection2(nums1: IntArray, nums2: IntArray): IntArray {
        val present = BooleanArray(1001)
        for (num in nums1) present[num] = true

        val intersection = IntArray(minOf(nums1.size, nums2.size))
        var count = 0
        for (num in nums2) {
            if (present[num]) {
                intersection[count++] = num
                present[num] = false
            }
        }
        return intersection.copyOf(count)        
    }

}
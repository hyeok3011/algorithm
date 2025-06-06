// https://leetcode.com/problems/container-with-most-water/description/?envType=study-plan-v2&envId=leetcode-75
class Solution {
    fun maxArea(height: IntArray): Int {
        var mostWater = 0  
        var left = 0
        var right = height.size - 1
        while (left < right) {
            val width = right - left
            val minHeight = minOf(height[left], height[right])
            mostWater = maxOf(mostWater, width * minHeight)
            if (height[left] > height[right]) {
                right -= 1
            } else {
                left += 1
            }
        }
        return mostWater
    }
}
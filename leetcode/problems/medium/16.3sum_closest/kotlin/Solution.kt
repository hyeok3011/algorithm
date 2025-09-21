class Solution {
    fun threeSumClosest(nums: IntArray, target: Int): Int {
        var answer = 0
        var diff = Int.MAX_VALUE
        nums.sort()
        for(i in 0 until nums.size - 2) {
            if (i != 0 && nums[i - 1] == nums[i]) continue
            var left = i + 1
            var right = nums.size - 1
            var sum: Int
            var currentDiff: Int
            while (left < right) {
                sum = nums[i] + nums[left] + nums[right]
                when {
                    sum > target -> right--
                    sum < target -> left++
                    sum == target -> return sum
                }
                currentDiff = abs(sum - target)
                if (diff > currentDiff) {
                    diff = currentDiff
                    answer = sum
                }
            }
        }
        return answer
    }

    // 처음부터 불가능한 부분 미리 처리하는 다른사람으 ㅣ코드.
    fun threeSumClosest2(nums: IntArray, target: Int): Int {
        Arrays.sort(nums)
        var closest = nums[0] + nums[1] + nums[2]
        if(closest == target) return target

        for(i in 0 until nums.size - 2) {
            if(i>0 && nums[i] == nums[i-1]) continue
            var j = i + 1
            var k = nums.size - 1
            var sum = nums[i] + nums[k] + nums[k-1] 
            if (sum < target) {
                closest = sum            
                continue
            }
            sum = nums[i] + nums[i+1] + nums[i+2]
            if (sum > target) {
                if (abs(sum-target) < abs(closest-target)) {
                    closest = sum
                }
                continue
            }
            while (j < k) {
                sum = nums[i] + nums[j] + nums[k]
                if (sum == target) {
                    return target
                } else if (sum > target) {
                    k--
                    while(k > j && nums[k] == nums[k+1]) {
                        k--
                    }
                } else {
                    j++
                    while(j<k && nums[j] == nums[j-1]) {
                        j++
                    }
                }
                if (abs(sum-target) < abs(closest-target)) {
                    closest = sum
                }
            }            
        }
        return closest
    }
}

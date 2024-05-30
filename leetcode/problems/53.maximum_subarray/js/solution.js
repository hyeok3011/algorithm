// https://leetcode.com/problems/maximum-subarray/
var maxSubArray = function(nums) {
    let maximumSum = nums[0];
    const dfs = (n) => {
        if (n === 0) {
            return nums[n];
        }
        let sum = Math.max(nums[n], dfs(n-1) + nums[n])
        maximumSum = Math.max(maximumSum, sum)
        return sum
    }
    dfs(nums.length - 1)
    return maximumSum; 
};

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let hashTable = new Map();
    let answer = new Array(2);
    for (let i = 0; i < nums.length; i++) {
        if (hashTable.has(target - nums[i])) {
            answer[0] = hashTable.get(target - nums[i]);
            answer[1] = i;
            break;
        } else {
            hashTable.set(nums[i], i);
        }
    }
    return answer;
};
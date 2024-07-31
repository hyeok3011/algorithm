/*
https://leetcode.com/problems/sort-colors/
문제 제약사항은 sort기능을 쓰지 않는것이고 처음에는 버블방식 시간복잡도는 O(n**2)로 풀었다.

다른사람들의 솔루션을 보니 네덜란드 국기 알고리즘(DNF)이라는 방식을 보고 다시 풀어봤다.
확실히 시간복잡도는 나으 ㅣ무식한 버블방식보다 더 효율적인듯하다. O(N)나오지 않을까 싶다.
*/

/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function(nums) {
    let low = 0, mid = 0, high = nums.length - 1;

    while (mid <= high) {
        if (nums[mid] === 0) {
            [nums[low], nums[mid]] = [nums[mid], nums[low]];
            low++;
            mid++;
        } else if (nums[mid] === 1) {
            mid++;
        } else {
            [nums[mid], nums[high]] = [nums[high], nums[mid]];
            high--;
        }
    }
    return nums;
};
// https://leetcode.com/problems/maximum-swap/description/
// O(N)방법
/**
 * @param {number} num
 * @return {number}
 */
var maximumSwap = function(num) {
    let numSplite = num.toString().split("");
    let maxIndex = -1, maxVal = -1, leftIndex = -1, swapIndex = -1;

    for (let i = numSplite.length - 1; i >= 0; i--) {
        let digit = parseInt(numSplite[i]);

        if (digit > maxVal) {
            maxVal = digit;
            maxIndex = i;
        } else if (digit < maxVal) {
            leftIndex = i;
            swapIndex = maxIndex;
        }
    }

    [numSplite[leftIndex], numSplite[swapIndex]] = [numSplite[swapIndex],numSplite[leftIndex]];
    return parseInt(numSplite.join(""));
};

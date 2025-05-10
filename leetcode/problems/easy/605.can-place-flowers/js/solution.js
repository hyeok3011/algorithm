/**
 * @param {number[]} flowerbed
 * @param {number} n
 * @return {boolean}
 */
var canPlaceFlowers = function(flowerbed, n) {
    for (let i = 0; i < flowerbed.length; i++) {
        if (flowerbed[i] === 1) {
            continue;
        }
        let preIndex = i - 1 <= 0 ? 0 : i - 1;
        let nextIndex = i + 1 >= flowerbed.length ? flowerbed.length - 1 : i + 1;
        if (flowerbed[preIndex] === 0 && flowerbed[nextIndex] === 0) {
            flowerbed[i] = 1;
            n--;
        }
    }

    return n <= 0;
};

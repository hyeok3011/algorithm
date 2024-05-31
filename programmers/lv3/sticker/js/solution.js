// https://school.programmers.co.kr/learn/courses/30/lessons/12971
// 어려운듯 어렵지 않은 문제였다
// 먼저 완전탐색은 불가능하다고 판단 내렸다. 2 ** 100000? 이나 되기 때문이다. 
// 가장 처음에 고민됐던 부분은 첫번째를 선택하는 경우 마지막을 탐색하지 않게 만드는것이였다.
// DP문제는 항상 row는 대부분 시간의 순서에 따르지만 col기준점을 세우기가 헷갈린다.
// 다른사람들의 풀이를 보니 마지막 return에서 dp의 마지막 인덱스가지고만 확인하는데 그래도 되는지 살짝 헷갈린다.
// 분명 마지막 이전 몇개에 더 높은 숫자가 있지 않을까라는 생각을 하는데 아직 제대로 이해하지 못한것인지 스스로도 어지럽다.

function solution(sticker) {
    if (sticker.length < 4) {
        return Math.max(...sticker);
    }
    
    let firstStartDP = new Array(sticker.length).fill(0);
    let secondStartDP = new Array(sticker.length).fill(0);
    firstStartDP[0] = sticker[0];
    secondStartDP[1] = sticker[1];
    
    for (let i = 2; i < sticker.length; i++) {
        if (i !== sticker.length -1) {
            firstStartDP[i] = Math.max(firstStartDP[i - 2] + sticker[i], (firstStartDP[i-3] || 0) + sticker[i]);
        }
        secondStartDP[i] = Math.max(secondStartDP[i - 2] + sticker[i], (secondStartDP[i-3] || 0) + sticker[i]);
    }

    let maxVal = Math.max(firstStartDP[sticker.length - 2], firstStartDP[sticker.length - 3])
    return Math.max(maxVal, secondStartDP[sticker.length - 1], secondStartDP[sticker.length - 2]);
}


console.log(solution([14, 6, 5, 11, 3, 9, 2, 10]))
console.log(solution([1, 3, 2, 5, 4]))
console.log(solution([4, 3]))
console.log(solution([4]))

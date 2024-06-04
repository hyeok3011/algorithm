/*
https://school.programmers.co.kr/learn/courses/30/lessons/42897
먼저 완전 탐색으로 풀경우 2^1000000이 되기때문에 빠르게 포기하자
그러면 DP를 활용하는 문제인듯하다.
최대 합을 구하는 문제이며 조건은 어레이에서 한 인덱스를 선택했을때 양옆의 인덱스를 선택하지 못한다는 것이다.
그 말은 즉 첫번째 인덱스를 선택한 경우 마지막 인덱스를 선택하지 못한다.
이와 매우 유사한 문제를 본적이 있을것이다.
바로 스티커 문제이다.
첫번째 시작 인덱스와 두번째 시작 인덱스를 따로 나누는 이유는 첫번째를 선택한 경우에는
마지막 인덱스를 사용하기 못하기 때문이다.
그 외에는 간단하다 점화식은 2개이다.
dp[i] = max(dp[i - 2] , dp[i - 3]) + num
dp[i] = max(dp[i-1], dp[i-2] + num) 이다.
*/ 
function solution(money) {
    const moneyLen = money.length;
    let firstStartDP = new Array(moneyLen).fill(0);
    let secondStartDP = new Array(moneyLen).fill(0);
    firstStartDP[0] = money[0];
    for (let i = 1; i < moneyLen; i++) {
        if (i !== moneyLen - 1) {
            firstStartDP[i] = Math.max(firstStartDP[i - 1], (firstStartDP[i - 2] || 0 )+ money[i])
        }
        secondStartDP[i] = Math.max(secondStartDP[i-1], secondStartDP[i - 2] || 0 + money[i])
    }

    return Math.max(firstStartDP[moneyLen - 2], secondStartDP[moneyLen - 1]);
}

// https://school.programmers.co.kr/learn/courses/30/lessons/43105
// triangle의 값은 500까지 나올 수 있다. 이말은 즉 500!이 된다는 건데 슈퍼컴퓨터정도는 필요하지 않을까 싶다.
// 어렵지 않은 DP문제이다.
function solution(triangle) {
    let dp = new Array(triangle.length).fill(null).map( (v, i) => new Array(i+1));
    
    let maxVal = 0;
    for (let i = 1; i < triangle.length; i++) {
        for (let j = 0; j < triangle[i].length; j++) {
            let largeVal = Math.max(dp[i-1][j - 1] || 0, dp[i-1][j] || 0)
            dp[i][j] = largeVal + triangle[i][j];
            maxVal = Math.max(maxVal, dp[i][j])
        }
    }
    
    return maxVal;
}

console.log(solution([[7], [3, 8], [8, 1, 0], [2, 7, 4, 4], [4, 5, 2, 6, 5]]))
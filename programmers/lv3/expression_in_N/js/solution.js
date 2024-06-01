// https://school.programmers.co.kr/learn/courses/30/lessons/42895
function solution(N, number) {
    if (N === number) {
        return 1;
    }

    let dp = new Array(9).fill(null).map(() => new Set());
    dp[1].add(N);
    let past = N;
    for (let i = 2; i <= 8; i++) {
        past = past * 10 + N;
        dp[i].add(past)
        for (let j = 1; j < i; j++) {
            dp[j].forEach(v => {
                dp[i-j].forEach(vv => {
                    dp[i].add(v * vv);
                    dp[i].add(Math.floor(v / vv));
                    dp[i].add(v + vv);
                    if (vv !== 0) {
                        dp[i].add(v - vv);
                    }
                })
            })
        }
        if (dp[i].has(number)) {
            return i
        }
    }
    return -1;
}

console.log(solution(5, 12))
console.log(solution(2, 11))
console.log(solution(1, 121)) // 4
console.log(solution(5, 3025)) // 4
console.log(solution( 5, 3125)) // 5
console.log(solution( 9, 1)) // 2
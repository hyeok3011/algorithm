function solution2(temperature, t1, t2, a, b, onboard) {
    if (temperature < t1) {
        temperature = t2 + t1 - temperature; // 편하게 오른쪽으로
    }
    temperature += 10;
    t1 += 10;
    t2 += 10;
    
    // row = minut col = 0 ~ 50 temperature
    let dp = new Array(onboard.length).fill(null).map(()=> new Array(51));
    dp.forEach(row => {
        row.fill(Infinity);
    })
    dp[0][temperature] = 0;
    for (let i = 1; i < onboard.length; i++) {
        let [start, end] = [t1, temperature];
        if (onboard[i] === 1) {
            [start, end] = [t1, t2]
        }
        for(let j = start; j <= end; j++) {
            if (j === temperature) {
                dp[i][j] = Math.min(dp[i-1][j-1], dp[i-1][j]);
            }else {
                dp[i][j] = Math.min(dp[i-1][j-1], dp[i-1][j] + b, dp[i-1][j+1] + a);
            }

        }
    }
    return Math.min(...dp[onboard.length - 1]);
}

function solution(temperature, t1, t2, a, b, onboard) {
    if (temperature < t1) {
        temperature = t2 + t1 - temperature; // 편하게 오른쪽으로
    }
    temperature += 10;
    t1 += 10;
    t2 += 10;
    
    // row = minut col = 0 ~ 50 temperature
    let dp = new Array(onboard.length).fill(null).map(()=> new Array(51));
    dp.forEach(row => {
        row.fill(Infinity);
    })
    dp[0][temperature] = 0;
    for (let i = 1; i < onboard.length; i++) {
        let [start, end] = [t1, temperature];
        if (onboard[i] === 1) {
            [start, end] = [t1, t2]
        }
        for(let j = start; j <= end; j++) {
            if (j === temperature) {
                dp[i][j] = Math.min(dp[i-1][j-1], dp[i-1][j]);
            }else {
                dp[i][j] = Math.min(dp[i-1][j-1], dp[i-1][j] + b, dp[i-1][j+1] + a);
            }

        }
    }
    return Math.min(...dp[onboard.length - 1]);
}


// console.log(solution(28, 18, 26, 10, 8, [0, 0, 1, 1, 1, 1, 1]))
console.log(solution(28, 18, 26, 10, 8, [0, 0, 1, 1, 1, 1, 1,0,0,1]))
// console.log(solution(11, 8, 10, 10, 100, [0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1]))
// console.log(solution(9, 7, 8, 10, 100, [0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1]))

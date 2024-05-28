function solution(k, n, reqs) {
    let requestByType = new Array(k).fill(null).map(() => []);
    reqs.forEach(request => {
        requestByType[request[2] - 1].push([request[0], request[1]]);
    })

    let waitTimeByMentorCount = new Array(k).fill(null).map(() => []);
    for (let i = 0; i < k; i++) {
        for (let j = 0; j <= n - k; j++) {
            waitTimeByMentorCount[i].push(calculateWaitTime(j+1, requestByType[i]));
        }
    }

    n = n - k;
    while( n > 0) {
        let maxWaitVal = 0;
        let maxWaitIndex = 0;
        waitTimeByMentorCount.forEach((waitTimes, type) => {
            let waitTime = waitTimes[0] - waitTimes[1];
            if (maxWaitVal < waitTime) {
                maxWaitVal = waitTime;
                maxWaitIndex = type;
                
            }
        })
        waitTimeByMentorCount[maxWaitIndex].shift();
        n--;
    }

    let totalWaitTIme = 0;

    waitTimeByMentorCount.forEach(waitTime => {
        totalWaitTIme += waitTime[0];
    })
    return totalWaitTIme
}

function calculateWaitTime(n, requests) {
    let waitTime = 0;
    let endTime = new Array(n).fill(0);
    requests = requests || [];

    requests.forEach(request => {
        let minIndex = 0;
        let minVal = 9999999;
        endTime.forEach((end, index) => {
            if (end < minVal) {
                minIndex = index;
                minVal = end;
            }
        })
        let diff = 0
        if (endTime[minIndex] > request[0]) {
            diff = endTime[minIndex] - request[0];
        }
        waitTime += diff;
        endTime[minIndex] = request[0] + request[1] + diff;
    })
    return waitTime;
}

console.log(solution(3, 5, [[10, 60, 1], [15, 100, 3], [20, 30, 1], [30, 50, 3], [50, 40, 1], [60, 30, 2], [65, 30, 1], [70, 100, 2]]))
console.log(solution(2, 3, [[5, 55, 2], [10, 90, 2], [20, 40, 2], [50, 45, 2], [100, 50, 2]]))

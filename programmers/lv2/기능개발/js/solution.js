/*
https://school.programmers.co.kr/learn/courses/30/lessons/42586?language=javascript
*/ 

function solution(progresses, speeds) {
    var answer = [];
    var day = 1;
    var releaseIndex = 0;
    var releaseCount = 0;
    while(releaseIndex < progresses.length) {
        releaseCount = 0;
        for (i = releaseIndex; i < progresses.length; i++) {
            if (day * (speeds[i]) + progresses[i] < 100) {
                break
            }
            releaseIndex += 1;             
            releaseCount += 1;
        }
        if (releaseCount != 0) {
            answer.push(releaseCount);
        }
        day += 1;
    }

    return answer;
}
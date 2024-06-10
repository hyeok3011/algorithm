/*
https://school.programmers.co.kr/learn/courses/30/lessons/148652
어.. 무슨 문제라고 해야하나....
그냥 패턴을 찾아야하는..? 문제이다.
결론만 말하자면 어떤 인덱스의 값이든 최초 어디에서 시작했는지를 파악하면 된다.
최오에 3번으로 시작하게되면 그 이후에가서도 0이며 최초 1번이여도 11011이기 때문에 0이 된다.
*/ 

function solution(n, l, r) {
    let answer = 0;
    const canto = [null, 1,1,0,1,1];
    for (let i = l; i <= r; i++) {
        let quotient = i
        let remainder = 0
        while (remainder != 3 && quotient > 5) {
            remainder = quotient % 5;
            quotient = Math.floor(quotient / 5)
            if (remainder != 0) {
                quotient++
            }
        }
        if (quotient != 3 && remainder != 3){
            answer++
        }
    }
    return answer;
}

// console.log(mo(108))
console.log(solution(2, 51, 75))
console.log(solution(2, 4, 17))
// 2	4	17
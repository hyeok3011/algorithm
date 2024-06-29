/*
https://school.programmers.co.kr/learn/courses/30/lessons/250135
개인적으로 좀 어려운 문제였다.
여러웠던 부분은 분침이나 시침이 초와 분에 따라서 조금씩 움직이는 메커니즘이 들어가서 더 어려웠다.
각도를 계산하며 여러 테스트를 진행다가 다른사람의 팁을 보고 나서야 풀 수 있었다.
시계는 1분에 최대 2번 알람이 울릴 수 있다. 
- 시/분 과 겹치는 경우
위 기준으로 00:00:00 ~ 23:59:59 까지는 2880번 울릴거 같지만 그렇지 않다.
문제 케이스에서 나와있듯이 2852번만 울린다고 한다. 
언제 울리지 않는지 보자.
1. 예외에 나와있듯이 59분 ~ 00분으로 넘어갈때 분/초 침이 만나지 않는듯 하다. = -24 (2880 - 24 = 2856)
2. 11시 59분에서 12:00될때 시/분/초 만나지 않는다. =  -2 (2856 - 2 = 2854)
3. 12:00:00은 시/분/초 가 모두 만나기 때문에 2번이 아닌 1번만 만난다. = -2 (2856 - 2 = 2854) 
그러면 이제 처음 시작 시간부터 목표 시간까지 계산한 뒤에 서로 빼버리면 그만이겠지만 한가지 더 문제가 있다.
시간에 초, 분, 시간의 위치이다. 위치에 따라서 알람의 개수가 바뀌게 된다.
침들은 원모양으로 움직이며 최대 360도까지 존재한다.
1. 초 = 1초에 6도씩 움직인다. 다른 분, 시간에 영향을 받지 않는다.
2. 분 = 분또한 1분에 6도씩 움직이나 초의 영향을 받게 된다. 분은 기본적으로 6도이나 초가 60번 움직일때 조금식 나눠서 움직이기 때문에
6 / 60 하면 0.1 이다.
3. 시 = 시는 분/초의 영향을 받는다. 기본적으로 한번에 30도씩 움직인다.
분에 영향은 시간은 단위는 위에서 말했듯이 30 / (분의 단위) 이다.
초의 영향은 분과 같은 공식에다가 / (초의 단위) 까지만 하면 된다.
위와 같은 방법으로 시 분 초 침의 각도 계산을 끝낼 수 있다.
*/

// 참고 https://school.programmers.co.kr/questions/63464
function solution(h1, m1, s1, h2, m2, s2) {
    const calculateAlarmCount = (hour,minute,second) => {
        let count = 0;
        count += (hour * 60 + minute) * 2;
        count -= hour;
        count -= 1;
        
        
        const hourDegree = ((hour % 12 * 30) + minute * 30 / 60 + second * 30 / 60 / 60) % 360
        const minuteDegree = ((minute * 6) + (second * 6 / 60)) % 360
        const secondDegree = second*6;

        if(secondDegree>=minuteDegree){
            count +=1
        }
        if(secondDegree>=hourDegree){
            count +=1
        }

        
        if(hour>=12){
            count -= 2
        } 
        return count
    }

    let alarmCount = calculateAlarmCount(h2,m2,s2) - calculateAlarmCount(h1,m1,s1)
    if ((h1 === 0|| h1 === 12)&& m1 === 0 && s1 === 0){
        alarmCount +=1
    }
    return alarmCount
}
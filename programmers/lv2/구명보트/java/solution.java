/*
https://school.programmers.co.kr/learn/courses/30/lessons/42885
작은것을 기준으로 큰것을 찾으려고하면 코드가 복잡해지게 되지만 큰것을 기준으로 작은것을 찾으면 코드가 훨 씬 간결해 진다.
*/ 
import java.util.*; 

class Solution {
    public int solution(int[] people, int limit) {
        int left = 0;
        int right = people.length - 1;
        int answer = 0;
        Arrays.sort(people);
        
        while (left <= right) {
            if (people[right] + people[left] <= limit) {
                left++;
            } 
            right--;
            answer++;
        }
        return answer;
    }
}
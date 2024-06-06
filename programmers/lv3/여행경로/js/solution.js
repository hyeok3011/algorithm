/*
https://school.programmers.co.kr/learn/courses/30/lessons/43164
*/ 

function solution(tickets) {
    let graph = {};

    tickets.forEach(v => {
    let [from, to] = [v[0], v[1]];
        if (!graph[from]) {
            graph[from] = []
        }
        graph[from].push(to);
    })

    for (const from in graph) {
        graph[from].sort();
    }

    const route = [];
    const stack = ["ICN"];

    while (stack.length) {
        const from = stack[stack.length - 1];
        if (graph[from] && graph[from].length > 0) {
            stack.push(graph[from].shift());
        } else {
            route.push(stack.pop());
        }
    }

    return route.reverse();
}


console.log(solution2([["ICN", "AAA"], ["ICN", "BBB"], ["BBB", "ICN"]]))
console.log(solution2([["ICN", "JFK"], ["HND", "IAD"], ["JFK", "HND"]]))
console.log(solution2([["ICN", "SFO"], ["ICN", "ATL"], ["SFO", "ATL"], ["ATL", "ICN"], ["ATL","SFO"]]))

console.log(solution2([["ICN", "D"], ["D", "ICN"], ["ICN", "B"]]))

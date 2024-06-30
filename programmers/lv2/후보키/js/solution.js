/*
https://school.programmers.co.kr/learn/courses/30/lessons/42890
어렵지는 않았지만 구현하는데 귀찮음이 있었다.
다른사람의 경우 비트연산을 푼코드도 있는데 좋은듯하다. @@@@@@@@@@@@@
*/

function solution(relation) {
    const numCols = relation[0].length;

    const generateCombination = (maximum, start, comLen, current, combinations) => {
        if (current.length === comLen) {
            combinations.push([...current]);
            return combinations;
        }
        for (let i = start; i < maximum - (comLen - current.length - 1); i++) {
            current.push(i);
            generateCombination(maximum, i + 1, comLen, current, combinations);
            current.pop();
        }
        return combinations;
    };

    const isCandidateKey = (combi, relation) => {
        let keySet = new Set();
        for (const row of relation) {
            let key = combi.map(index => row[index]).join('|');
            if (keySet.has(key)) return false;
            keySet.add(key);
        }
        return true;
    };

    const isMinimal = (combi, candidateKeys) => {
        for (const key of candidateKeys) {
            if (key.every(element => combi.includes(element))) return false;
        }
        return true;
    };

    let answer = 0;
    let candidateKeys = [];

    for (let i = 1; i <= numCols; i++) {
        for (const combi of generateCombination(numCols, 0, i, [], [])) {
            if (isCandidateKey(combi, relation) && isMinimal(combi, candidateKeys)) {
                candidateKeys.push(combi);
                answer++;
            }
        }
    }

    return answer;
}

// console.log(solution([["100","ryan","music","2"],["200","apeach","math","2"],["300","tube","computer","3"],["400","con","computer","4"],["500","muzi","music","3"],["600","apeach","music","2"]]))
console.log(solution([["a","1","aaa","c","ng"],
["a","1","bbb","e","g"],
["c","1","aaa","d","ng"],
["d","2","bbb","d","ng"]]))

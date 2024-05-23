// https://school.programmers.co.kr/learn/courses/30/lessons/250136#qna
/*
dfs로 풀어보니 이상하게 효율성 테스트 하나가 통과되지 않는다. 아마도 콜스택의 문제인듯 하다.
bfs로 풀어보니 문제없이 동작한다.
*/

function solution(land) {
    let m = land.length;
    let n = land[0].length;
    let colTotalOil = new Array(n).fill(0);
    
    let queue = new Array();

    for (let col = 0; col < n; col++) {
        for (let row = 0; row < m; row++) {
            if (land[row][col] === 0) {
                continue;
            }
            queue.push([row, col])
            let size = 0;
            let minCol = col;
            let maxCol = col;
            while(queue.length > 0) {
                let [currentRow, currentCol] = queue.shift();
                if (currentRow < 0 || currentCol < 0 || currentRow >= m || currentCol >= n) {
                    continue;
                }

                if (land[currentRow][currentCol] === 0) {
                    continue;
                }

                size++;
                land[currentRow][currentCol] = 0;

                minCol = Math.min(minCol, currentCol);
                maxCol = Math.max(maxCol, currentCol);

                queue.push([currentRow + 1, currentCol]);
                queue.push([currentRow - 1, currentCol]);
                queue.push([currentRow, currentCol + 1]);
                queue.push([currentRow, currentCol - 1]);
            }

            for (let i = minCol; i <= maxCol; i++) {
                colTotalOil[i] += size;
            }
        }
    }

    return Math.max(...colTotalOil)
}

function solutionDFS(land) {
    let m = land.length;
    let n = land[0].length;
    let colTotalOil = new Array(n).fill(0);

    const recursion = (row, col, size, minMax) => {
        land[row][col] = 0;
        minMax.min = Math.min(minMax.min, col)
        minMax.max = Math.max(minMax.max, col)
        size++
        if (row - 1 > 0 && land[row -1][col]) {
            size = recursion(row - 1, col, size, minMax);
        }

        if (row + 1 < m && land[row + 1][col]) {
            size = recursion(row + 1, col, size, minMax);
        }

        if (col - 1 > 0 && land[row][col - 1]) {
            size = recursion(row, col - 1, size, minMax);
        }

        if (col + 1 < n && land[row][col + 1]) {
            size = recursion(row, col + 1, size, minMax);
        }

        return size;
    }

    for (let col = 0; col < n; col++) {
        for (let row = 0; row < m; row++) {
            if (land[row][col]) {
                let minMax = {min: col, max: col}
                let size = recursion(row, col, 0, minMax);
                for (let i = minMax.min; i <= minMax.max; i++) {
                    colTotalOil[i] += size;
                }
            }
        }
    }
    return Math.max(...colTotalOil);
}





solution([[0, 0, 0, 1, 1, 1, 0, 0], [0, 0, 0, 0, 1, 1, 0, 0], [1, 1, 0, 0, 0, 1, 1, 0], [1, 1, 1, 0, 0, 0, 0, 0], [1, 1, 1, 0, 0, 0, 1, 1]])
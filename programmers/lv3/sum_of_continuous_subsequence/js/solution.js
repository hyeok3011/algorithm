
function solution2(sequence) {
    let negativeStart = 0;
    let positiveStart = 0;
    let maxPositiveStart = -100000;
    let maxNegativeStart = -100000;

    sequence.forEach((v, i) => {
        if (i % 2 == 0) {
            positiveStart = Math.max(positiveStart + v, v);
            negativeStart = Math.max(negativeStart + -v, -v);
        } else {
            positiveStart = Math.max(positiveStart + -v, -v);
            negativeStart = Math.max(negativeStart + v, v);
        }
        maxPositiveStart = Math.max(maxPositiveStart, positiveStart);
        maxNegativeStart = Math.max(maxNegativeStart, negativeStart);
    })
    
    return Math.max(maxPositiveStart, maxNegativeStart);
}
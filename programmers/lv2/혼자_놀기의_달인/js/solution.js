/*
https://school.programmers.co.kr/learn/courses/30/lessons/131130
*/ 

function solution(cards) {
    let choiceCardSet = new Set();
    let boxGroupCount = new Array(cards.length);
    
    for (let i = 0; i < cards.length; i++) {
        if (choiceCardSet.has(cards[i])) {
            boxGroupCount[i] = 0
        }
        
        let boxGroup = new Set();
        let currentCard = cards[i];
        while (!boxGroup.has(currentCard) && !choiceCardSet.has(currentCard)) {
            boxGroup.add(currentCard);
            choiceCardSet.add(currentCard);
            currentCard = cards[currentCard - 1];
        }
        boxGroupCount[i] = boxGroup.size;
    }
    
    boxGroupCount.sort((a,b) => b - a)
    return boxGroupCount[0] * boxGroupCount[1];
}
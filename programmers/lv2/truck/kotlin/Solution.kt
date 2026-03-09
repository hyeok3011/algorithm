// https://school.programmers.co.kr/learn/courses/30/lessons/42583?language=kotlin
class Solution {
    fun solution(bridgeLength: Int, weight: Int, truckWeights: IntArray): Int {
        var answer = 0
        val truckQueue = ArrayDeque<Int>()
        val timeQueue = ArrayDeque<Int>()
        var bridgeWeightSum = 0
        for (truckWeight in truckWeights) {
            while (bridgeWeightSum + truckWeight > weight){
                val nextExit = timeQueue.removeFirst()
                answer = maxOf(answer, nextExit)
                bridgeWeightSum -= truckQueue.removeFirst()
            }
            timeQueue.add(answer + bridgeLength)
            truckQueue.add(truckWeight)
            bridgeWeightSum += truckWeight
            answer++
        }
        return answer + bridgeLength
    }
}
object EqualSidesOfAnArray {
    fun findEvenIndex(arr: IntArray): Int {
        val totalSum = arr.sum()
        var leftSum = 0
        arr.forEachIndexed { index, value ->
            val rightSum = totalSum - leftSum - value
            if (leftSum == rightSum) {
                return index
            }
            leftSum += value
        }
        return -1
    }
}

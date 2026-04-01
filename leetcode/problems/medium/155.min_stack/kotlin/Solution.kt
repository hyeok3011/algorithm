// https://leetcode.com/problems/min-stack/description/?envType=problem-list-v2&envId=rab78cw1
// LinkedList로 직접 구현한 사람들도 있음...ㄷㄷ
class MinStack() {
    private class Item(val value: Int, val minValue: Int )
    private val stack = ArrayDeque<Item>()
    fun push(`val`: Int) {
        val minVal = if (stack.isEmpty()) `val` else minOf(getMin(), `val`)
        stack.addLast(Item(`val`, minVal))
    }

    fun pop() {
        stack.removeLast()
    }

    fun top(): Int = stack.last().value

    fun getMin(): Int = stack.last().minValue
}

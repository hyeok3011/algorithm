// https://leetcode.com/problems/implement-queue-using-stacks/

class MyQueue {
    private val inStack = ArrayDeque<Int>()
    private val outStack = ArrayDeque<Int>()

    fun push(x: Int) {
        inStack.add(x)
    }

    fun pop(): Int {
        if (outStack.isEmpty()) {
            transferInToOut()
        }
        return outStack.removeLast()
    }

    fun peek(): Int {
        if (outStack.isEmpty()) {
            transferInToOut()
        }
        return outStack.last()
    }

    fun empty(): Boolean = inStack.isEmpty() && outStack.isEmpty()

    private fun transferInToOut() {
        while (inStack.isNotEmpty()) {
            outStack.add(inStack.removeLast())
        }
    }
}

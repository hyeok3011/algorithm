// https://leetcode.com/problems/add-two-numbers/

class Solution {
    fun addTwoNumbers(l1: ListNode?, l2: ListNode?): ListNode? {
        return addWithCarry(l1, l2, 0)
    }

    private fun addWithCarry(l1: ListNode?, l2: ListNode?, overflow: Int): ListNode? {
        if (l1 == null && l2 == null && overflow == 0) {
            return null
        }

        var sum: Int = overflow
        sum += l1?.`val` ?: 0
        sum += l2?.`val` ?: 0
        return currentNode(sum % 10).apply {
            next = addWithCarry(l1?.next, l2?.next, sum / 10)
        }
    }
}

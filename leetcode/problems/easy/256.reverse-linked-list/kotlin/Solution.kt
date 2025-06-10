// https://leetcode.com/problems/reverse-linked-list/?envType=study-plan-v2&envId=leetcode-75
// 진짜 쉬운문제인데 오랜만에 스왑관련된 문제 푸니까 너무 어려웠음....
/**
 * Example:
 * var li = ListNode(5)
 * var v = li.`val`
 * Definition for singly-linked list.
 * class ListNode(var `val`: Int) {
 *     var next: ListNode? = null
 * }
 */
class Solution {
    fun reverseList(head: ListNode?): ListNode? {
        var pre: ListNode? = null
        var current = head

        while (current != null) {
            val temp = current.next
            current.next = pre
            pre = current
            current = temp
        }

        return pre
    }
}
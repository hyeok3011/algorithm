//  https://leetcode.com/problems/middle-of-the-linked-list/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun middleNode(head: ListNode?): ListNode? {
        var slow = head
        var fast = head
        while(fast?.next != null) {
            slow = slow?.next
            fast = fast.next.next
        }

        return slow
    }
}
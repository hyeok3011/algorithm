package main

// https://leetcode.com/problems/odd-even-linked-list/?envType=study-plan-v2&envId=leetcode-75
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    odd := head
    evenHead := head.Next
    even := evenHead
    for even != nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }
    odd.Next = evenHead
    return head
}

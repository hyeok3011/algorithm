package main

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/?envType=study-plan-v2&envId=leetcode-75
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }

    slow, fast := head, head.Next.Next
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }

    slow.Next = slow.Next.Next

    return head
}

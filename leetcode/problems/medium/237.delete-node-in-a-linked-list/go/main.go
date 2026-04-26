package main
//  https://leetcode.com/problems/delete-node-in-a-linked-list/
// ....? 뭐지..?
func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}

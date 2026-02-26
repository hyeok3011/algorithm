package main

// https://leetcode.com/problems/linked-list-cycle-ii/
// @@@@@ 처음에는 단순하게 set()을 사용하여 풀었지만 O(1)로 푸는거는 상당히 어려워 다른사람의 풀이를 보고 풀었음.
// https://leetcode.com/problems/linked-list-cycle-ii/solutions/3019604/go-easy-to-understand-on-time-o1-space-w-2ln3/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

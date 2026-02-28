package main

// https://leetcode.com/problems/insertion-sort-list/description/
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -5001,
	}

	for head != nil {
		current := dummy.Next
		pre := dummy
		for current != nil && current.Val < head.Val {
			pre = current
			current = current.Next
		}

		next := head.Next
		head.Next = current
		pre.Next = head
		head = next
	}

	return dummy.Next
}

// 다른사람들의 최적화 방법, tail에 next를 이어주고 Next의 next가 정렬되어있는 상태라면 bypass
// @@@@@@
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newListNode := &ListNode{Next: head}
	tail := head
	curr := head.Next

	for curr != nil {
		if curr.Val >= tail.Val {
			// 이미 연결되어있음.
			tail = curr
			curr = curr.Next
		} else {
			prev := newListNode
			for prev.Next != curr && prev.Next.Val <= curr.Val {
				prev = prev.Next
			}

			prevNxt := prev.Next
			currNxt := curr.Next

			prev.Next = curr
			curr.Next = prevNxt

			tail.Next = currNxt

			curr = currNxt
		}
	}

	return newListNode.Next
}

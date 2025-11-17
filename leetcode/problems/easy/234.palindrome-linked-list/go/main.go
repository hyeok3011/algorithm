// https://leetcode.com/problems/palindrome-linked-list/

func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev := slow
	current := slow
	next := slow.Next

	for next != nil {
		temp := next.Next
		current.Next = prev
		next.Next = current
		prev = current
		current = next
		next = temp
	}

	for current != head {
		if head.Val != current.Val {
			return false
		}
		head = head.Next
		current = current.Next
	}
	return true
}

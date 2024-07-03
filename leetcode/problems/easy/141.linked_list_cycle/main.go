package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.com/problems/linked-list-cycle/description/
해당 노드가 사이클을 만드는지 확인하는 문제이다.
매우 간단한 문제라고 생각하고 nodeSet을 만들어 풀었다.
하지만 메모리 퍼포먼스가 좋지 고민하다 다른사람의 풀이를 보니
two-pointer slow fast방식푼 아이디어를 얻어 다시 풀어봤다.
위 방식으로 풀면 확실히 리스트의 길이가 길어질수록 효율적일듯 하다

Runtime 0 ms Beats 100% Memory 4.5 MB Beats 65.39%
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodeTable := make(map[*ListNode]bool)
	currentNode := head
	for currentNode != nil {
		if nodeTable[currentNode] {
			return true
		}
		nodeTable[currentNode] = true
		currentNode = currentNode.Next
	}
	return false
}

func twoPonterSolution(head *ListNode) bool {
	slowPointer, fastPointer := head, head
	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}

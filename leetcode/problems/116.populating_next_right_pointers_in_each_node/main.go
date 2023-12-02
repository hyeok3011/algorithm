package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/
Node의 같은 레벨의 오른쪽 노드를 연결 시켜주는 문제이다.
생각보다 매우 간단한 문제이다. 레벨을 내려가며 같은 레벨의 노드들을 탐색해야하기 때문에 BFS를 사용해서 풀었다.
Runtime 0 ms Beats 100% Memory 6.6 MB Beats 56.34%
*/

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	var sameLevelNodeCount int
	var preNode, currentNode *Node
	queue := []*Node{root.Left, root.Right}
	for len(queue) > 0 {
		sameLevelNodeCount = len(queue)

		for i := 0; i < sameLevelNodeCount; i++ {
			if queue[i] == nil {
				continue
			}
			currentNode = queue[i]

			if preNode != nil {
				preNode.Next = currentNode
			}

			queue = append(queue, currentNode.Left, currentNode.Right)
			preNode = currentNode
		}
		queue = queue[sameLevelNodeCount:]
		preNode = nil
	}

	return root
}

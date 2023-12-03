package main

/*
https://leetcode.com/problems/lru-cache/
https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU
Least Recently used 알고리즘은 구현하는 문제이다.
LRU알고리즘을 redis에 쓰여 간략하게 공부한적은 있는데 실제로 구현해본적은 없다.

자료구조는 hashtable과 doublelinkedlist를 사용하였고
head와 tail에는 더미 노드를 사용했다.
처음에는 더미노드를 사용하지 않고 실제 노드를 스왑하는 방식을 사용했는데 if문 코드가 너무 더러워 변경했다.

Runtime 409 ms Beats 96.58% Memory 65.1 MB Beats 99.41%
*/

type Node struct {
	key, value     int
	previous, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.previous = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		lruNode := this.tail.previous
		delete(this.cache, lruNode.key)
		this.removeNode(lruNode)
	}

	newNode := &Node{key: key, value: value}
	this.cache[key] = newNode
	this.addToFront(newNode)
}

func (this *LRUCache) addToFront(node *Node) {
	node.previous = this.head
	node.next = this.head.next
	this.head.next.previous = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	if node == this.head || node == this.tail {
		return
	}

	previous := node.previous
	next := node.next
	previous.next = next
	next.previous = previous
}

func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addToFront(node)
}

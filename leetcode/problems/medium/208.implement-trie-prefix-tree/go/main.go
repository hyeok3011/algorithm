package main

// https://leetcode.com/problems/implement-trie-prefix-tree/?envType=study-plan-v2&envId=leetcode-75
type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{
		root: &trieNode{},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i := range word {
		wordIndex := word[i] - 'a'
		if node.children[wordIndex] == nil {
			node.children[wordIndex] = &trieNode{}
		}
		node = node.children[wordIndex]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.search(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.search(prefix)
	return node != nil
}

func (t *Trie) search(word string) *trieNode {
	node := t.root
	for i := range word {
		wordIndex := word[i] - 'a'
		if node.children[wordIndex] == nil {
			return nil
		}
		node = node.children[wordIndex]
	}
	return node
}

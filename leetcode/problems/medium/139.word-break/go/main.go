package main

// https://leetcode.com/problems/word-break/?envType=problem-list-v2&envId=rab78cw1
// @@@@@ 구현이 어려웠음... 다시 풀어볼것..
type trieNode struct {
    children [26]*trieNode
    isEnd bool
}

type trie struct {
    root *trieNode
}

func (t *trie) GetRoot() *trieNode {
    return t.root
}

func (t *trie) Insert(word string) {
    node := t.root
    for i := range word {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &trieNode{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func wordBreak(s string, wordDict []string) bool {
    t := trie{
        root: &trieNode{},
    }

    for _, v := range wordDict {
        t.Insert(v)
    }
    impassable := make(map[int]bool)
    var recursion func(int) bool
    recursion = func(startIndex int) bool {
        if impassable[startIndex]{
            return false
        }

        if startIndex >= len(s) {
            return true
        }
        node := t.GetRoot()
        for i := startIndex; i < len(s); i++ {
            idx := s[i] - 'a'
            if node.children[idx] == nil {
                break
            }
            node = node.children[idx]
            if node.isEnd {
                if recursion(i + 1) {
                    return true
                }
            }
        }

        impassable[startIndex] = true
        return false
    }

    return recursion(0)
}

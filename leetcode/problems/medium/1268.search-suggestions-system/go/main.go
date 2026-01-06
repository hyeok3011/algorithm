package main

// https://leetcode.com/problems/search-suggestions-system/?envType=study-plan-v2&envId=leetcode-75
// @@@@ 현재 캐시는 적용하지 않음. 캐시까지 적용해서 다시 풀어볼것
func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products) 

    t := NewTrie()
    for _, p := range products {
        t.Insert(p)
    }

    results := make([][]string, 0)
    
    for i := 0; i < len(searchWord); i++ {
        prefix := searchWord[:i+1]
        results = append(results, t.Search(prefix, 3))
    }

    return results
}

type trieNode struct {
    children [26]*trieNode
    isEnd    bool
}

type trie struct {
    root *trieNode
}

func NewTrie() *trie {
    return &trie{
        root: &trieNode{},
    }
}

func (t *trie) Insert(word string) {
    node := t.root
    for i := range word {
        index := word[i] - 'a'
        if node.children[index] == nil {
            node.children[index] = &trieNode{}
        }
        node = node.children[index]
    }
    node.isEnd = true
}

func (t *trie) Search(prefix string, limit int) []string {
    node := t.findNode(prefix)
    if node == nil {
        return []string{}
    }

    results := []string{}
    t.collectSuggestions(node, prefix, &results, limit)
    return results
}

func (t *trie) findNode(prefix string) *trieNode {
    node := t.root
    for i := range prefix {
        index := prefix[i] - 'a'
        if node.children[index] == nil {
            return nil
        }
        node = node.children[index]
    }
    return node
}

func (t *trie) collectSuggestions(node *trieNode, currentWord string, results *[]string, limit int) {
    if len(*results) == limit {
        return
    }

    if node.isEnd {
        *results = append(*results, currentWord)
    }

    for i := 0; i < 26; i++ {
        if node.children[i] != nil {
            nextChar := string('a' + rune(i))
            t.collectSuggestions(node.children[i], currentWord+nextChar, results, limit)
            
            if len(*results) == limit {
                return
            }
        }
    }
}

// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/

class WordDictionary() {
    val root = TrieNode()
    fun addWord(word: String) {
        var node = root
        for (char in word) {
            val charIndex = char - 'a'
            if (node.children[charIndex] == null) {
                node.children[charIndex] = TrieNode()
            }
            node = node.children[charIndex]!!
        }

        node.isEnd = true
    }

    fun search(word: String): Boolean {
        return search(word, 0, root)
    }

    fun search(word: String, index: Int, node: TrieNode): Boolean {
        if (index == word.length) {
            return node.isEnd
        }
        val char = word[index]
        if (char == '.') {
            for (childNode in node.children) {
                if (childNode == null) {
                    continue
                }
                if (search(word, index + 1, childNode)) {
                    return true
                }
            }
            return false
        } else {
            val charIndex = char - 'a'
            return node.children[charIndex]?.let { search(word, index + 1, it) } ?: false
        }
    }
    class TrieNode {
        val children = arrayOfNulls<TrieNode>(26)
        var isEnd: Boolean = false
    }
}
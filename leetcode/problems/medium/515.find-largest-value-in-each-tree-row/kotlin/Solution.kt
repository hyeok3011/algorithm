// https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/
class Solution {
    fun largestValues(root: TreeNode?): List<Int> {
        if (root == null) return emptyList()

        val answer = mutableListOf<Int>()
        val queue = ArrayDeque<TreeNode>()
        queue.addFirst(root)

        while (queue.isNotEmpty()) {
            var maxVal = Int.MIN_VALUE
            val level = queue.size
            repeat(level) {
                val node = queue.removeLast()
                maxVal = maxOf(maxVal, node.`val`)
                node.left?.let { queue.addFirst(it) }
                node.right?.let { queue.addFirst(it) }
            }
            answer.add(maxVal)
        }
        return answer
    }
}
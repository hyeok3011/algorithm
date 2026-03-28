// https://leetcode.com/problems/clone-graph/?envType=problem-list-v2&envId=rab78cw1
/**
 * Definition for a Node.
 * class Node(var `val`: Int) {
 *     var neighbors: ArrayList<Node?> = ArrayList<Node?>()
 * }
 */

 class Solution {
    fun cloneGraph(node: Node?): Node? {
        if (node == null) return null

        val visited = mutableMapOf<Int, Node?>()
        fun recursion(current: Node?): Node? {
            if (current == null) return null
            visited[current.`val`]?.let{return it}

            val newNode = Node(current.`val`)
            visited[newNode.`val`] = newNode
            for (neighbor in current.neighbors) {
                newNode.neighbors.add(recursion(neighbor))
            }
            return newNode
        }
        return recursion(node)
    }
}
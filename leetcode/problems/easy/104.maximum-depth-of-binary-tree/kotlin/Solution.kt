// https://leetcode.com/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&envId=leetcode-75

/**
 * Example:
 * var ti = TreeNode(5)
 * var v = ti.`val`
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */
class Solution {
    fun maxDepth(root: TreeNode?): Int {
        return dfs(root, 0)
    }

    fun dfs(node: TreeNode?, depth: Int): Int {
        if (node == null) {
            return depth
        }

        return max(dfs(node.left, depth + 1), dfs(node.right, depth + 1))
    }
}


// 이렇게도 정리가 가능함........
class Solution {
    fun maxDepth(root: TreeNode?): Int {
        if (root == null) return 0
        return 1 + max(maxDepth(root.left), maxDepth(root.right))
    }
}
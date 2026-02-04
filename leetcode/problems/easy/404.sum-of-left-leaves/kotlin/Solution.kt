// https://leetcode.com/problems/sum-of-left-leaves/
class Solution {
    fun sumOfLeftLeaves(root: TreeNode?): Int {
        return recursion(root, false)
    }

    fun recursion(node: TreeNode?, isLeft: Boolean): Int {
        if (node == null) return 0

        if (node.left == null && node.right == null && isLeft) {
            return node.`val`
        }

        return recursion(node.left, true) + recursion(node.right, false)
    }
}

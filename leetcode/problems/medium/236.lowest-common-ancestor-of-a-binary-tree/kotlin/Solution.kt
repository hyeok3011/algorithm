// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
// kotlin으로 다시 풀었지만 golang과 똑같은 풀이였음.
// 아래의 풀이는 다른사람의 풀이...
// @@@@
class Solution {
    fun lowestCommonAncestor(root: TreeNode?, p: TreeNode?, q: TreeNode?): TreeNode? {
        if (root == null || root === p || root === q) return root

        val left = lowestCommonAncestor(root.left, p, q)
        val right = lowestCommonAncestor(root.right, p, q)

        return when {
            left != null && right != null -> root
            else -> left ?: right                
        }
    }
}
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun kthSmallest(root: TreeNode?, k: Int): Int {
        var n = k
        var result = 0
        fun recursion(node: TreeNode?){
            if (node == null || n == 0) return
            recursion(node.left)
            n--
            if (n == 0) result = node.`val`
            recursion(node.right)
        }
        recursion(root)
        return result
    }
}

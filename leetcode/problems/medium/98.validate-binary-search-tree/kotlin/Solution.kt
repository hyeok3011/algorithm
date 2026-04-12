// https://leetcode.com/problems/validate-binary-search-tree/description/
class Solution {
    fun isValidBST(root: TreeNode?): Boolean {
		fun validBST(node: TreeNode?, min: Long, max: Long): Boolean {
			if (node == null) return true
			val nodeVal = node.`val`
			if (nodeVal <= min || max <= nodeVal) return false
				
			return validBST(node.left, min, minOf(nodeVal.toLong(), max)) &&
			validBST(node.right, minOf(max, nodeVal.toLong()), max)
		}
		return validBST(root, Long.MIN_VALUE, Long.MAX_VALUE)
    }
}

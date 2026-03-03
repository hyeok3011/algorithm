// https://leetcode.com/problems/binary-search-tree-iterator/
// O(h)를 확인하고 다시 풀이.
class BSTIterator(root: TreeNode?) {
    val stack = ArrayDeque<TreeNode>()

    init {
        recursion(root)
    }

    private fun recursion(node: TreeNode?) {
        if (node == null) return
        stack.add(node)
        recursion(node.left)
    }

    fun next(): Int {
        val node = stack.removeLast()
        recursion(node.right)
        return node.`val`
    }

    fun hasNext(): Boolean = stack.isNotEmpty()
}

class BSTIterator2(root: TreeNode?) {
    val queue = ArrayDeque<Int>()
 
     init {
         inorder(root)
     }
 
     private fun inorder(node: TreeNode?) {
         if (node == null) return
         
         inorder(node.left)    
         queue.add(node.`val`) 
         inorder(node.right)
     }
 
     fun next(): Int {
         if (queue.size == 0) {
             return -1
         }
         return queue.removeFirst()
     }
 
     fun hasNext(): Boolean = queue.isNotEmpty()
 }
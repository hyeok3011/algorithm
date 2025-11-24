// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    
    numbers := make([]int, 0)
    
    var recursion func(node *Node)
    recursion = func(node *Node) {
        if node == nil {
            return
        }
        numbers = append(numbers, node.Val)
        
        for _, child := range node.Children {
            recursion(child)
        }
    }
    
    recursion(root)
    return numbers
}

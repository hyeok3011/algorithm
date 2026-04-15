// https://leetcode.com/problems/n-ary-tree-postorder-traversal/
func postorder(root *Node) []int {
    result := make([]int, 0)
    var dfs func(*Node)
    dfs = func(node *Node) {
        if node == nil {
            return
        }
        for _, child := range node.Children {
            dfs(child)
        }
        result = append(result, node.Val)
    }
    dfs(root)
    return result
}

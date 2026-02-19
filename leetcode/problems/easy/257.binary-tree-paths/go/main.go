package main

// https://leetcode.com/problems/binary-tree-paths/description/
func binaryTreePaths(root *TreeNode) []string {
    paths := []string{}

    var dfs func (node *TreeNode, path string)
    dfs = func (node *TreeNode, path string){
        if node == nil{
            return
        }
        path += strconv.Itoa(node.Val)
        if node.Left == nil && node.Right == nil{
            paths = append(paths, path)
            return
        }
        path += "->"
        dfs(node.Left, path)
        dfs(node.Right, path)
    }
    dfs(root, "")

    return paths
}

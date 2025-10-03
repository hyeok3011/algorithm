// https://leetcode.com/problems/recover-binary-search-tree/description/?envType=problem-list-v2&envId=depth-first-search
// @@@@@@@@
// 또 풀어볼것 내가 제대로 이해하지 못하는 방식으로 다른사람들의 풀이가 많은듯함.
func recoverTree(root *TreeNode) {
    var firstWrongNode, secondWrongNode, pre *TreeNode
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode){
        if node == nil {
            return
        }
        dfs(node.Left)
        if pre != nil && pre.Val > node.Val {
            if firstWrongNode == nil{
                firstWrongNode = pre
            }
            if firstWrongNode != nil{
                secondWrongNode = node
            }
        }
        pre = node
        dfs(node.Right)
    }
    dfs(root)
    firstWrongNode.Val, secondWrongNode.Val = secondWrongNode.Val, firstWrongNode.Val
}


// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	path := make(map[int]bool)
    commonNodes := []*TreeNode{}
    var recursion func(node *TreeNode, targetNum int)
    recursion = func(node *TreeNode, targetNum int) {
        if path[node.Val] {
            commonNodes = append(commonNodes, node) 
        }
        path[node.Val] = true

        if node.Val == targetNum {
            return
        }

        if node.Val > targetNum {
            recursion(node.Left, targetNum)
        } else {
            recursion(node.Right, targetNum)
        }
    }
    recursion(root, p.Val)
    recursion(root, q.Val)
    return commonNodes[len(commonNodes) - 1]
}

// 다른 사람의 풀이
// root자체를 옮기는 방법, 생각하지 못했음. 
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
    curr := root

    for curr != nil {
        if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
        } else if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
        } else {
            return curr
        }
    }

    return nil
}

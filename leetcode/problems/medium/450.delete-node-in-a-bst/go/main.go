package main

// https://leetcode.com/problems/delete-node-in-a-bst/?envType=study-plan-v2&envId=leetcode-75

// Value Swap으로 수정...
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }

        minNode := findMin(root.Right)
        
        root.Val = minNode.Val
        
        root.Right = deleteNode(root.Right, minNode.Val)
    }

    return root
}


// 실제 노드 자체를 옮겨버리는 방식. > 트리의 height에 문제가 생기는 구조
func findMin(node *TreeNode) *TreeNode {
    if node.Left != nil {
        return findMin(node.Left)
    }
    return node
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val == key {
        if root.Right != nil {
            rightLeafNode := findMin(root.Right)
            rightLeafNode.Left = root.Left
            return root.Right
        } else  {
            return root.Left
        }
    } else if root.Val < key {
        root.Right = deleteNode(root.Right, key)
    } else {
        root.Left = deleteNode(root.Left, key)
    }

    return root
}

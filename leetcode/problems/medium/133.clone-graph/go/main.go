package main

// https://leetcode.com/problems/clone-graph/?envType=problem-list-v2&envId=rab78cw1
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    visited := make(map[int]*Node)
    var recursion func(node *Node) *Node
    recursion = func(node *Node) *Node {
        if existNode, ok := visited[node.Val]; ok {
            return existNode
        }

        newNode := &Node{
            Val: node.Val,
            Neighbors: []*Node{},
        }
        visited[node.Val] = newNode

        for i := range node.Neighbors {
            newNode.Neighbors = append(newNode.Neighbors, recursion(node.Neighbors[i]))
        }

        return newNode
    }

    return recursion(node)
}

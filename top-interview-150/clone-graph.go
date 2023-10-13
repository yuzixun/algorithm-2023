package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)
	var dfs func(node *Node) *Node

	dfs = func(cur *Node) *Node {

		visitedNode, exist := visited[cur.Val]
		if exist {
			return visitedNode
		}

		newNode := &Node{Val: cur.Val}
		visited[cur.Val] = newNode
		for _, neighbor := range cur.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor))
		}

		return newNode
	}

	return dfs(node)
}

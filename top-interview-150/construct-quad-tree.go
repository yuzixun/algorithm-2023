package main

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	return constructGrid(grid, 0, len(grid[0])-1, 0, len(grid)-1)
}

func constructGrid(grid [][]int, l, r, u, d int) *Node {

	if isLeaf(grid, l, r, u, d) {
		//return &Node{Val: grid[u][l] == 1, IsLeaf: true}
	}
	m, h := (r+l)/2, (d+u)/2

	node := new(Node)
	//node.Val = true
	node.IsLeaf = false
	node.TopLeft = constructGrid(grid, l, m, u, h)
	node.TopRight = constructGrid(grid, m+1, r, u, h)
	node.BottomLeft = constructGrid(grid, l, m, h+1, d)
	node.BottomRight = constructGrid(grid, m+1, r, h+1, d)
	return node
}

func isLeaf(grid [][]int, l, r, u, d int) bool {
	val := grid[u][l]
	for i := l; i <= r; i++ {
		for j := u; j <= d; j++ {
			if val != grid[j][i] {
				return false
			}
		}
	}
	return true
}
¬

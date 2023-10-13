package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node

	Left  *Node
	Right *Node

	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node

	Neighbors []*Node
}

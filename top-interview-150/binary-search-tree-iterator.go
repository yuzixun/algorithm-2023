package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{}
	bst.buildBST(root)
	return bst
}

func (this *BSTIterator) buildBST(root *TreeNode) {
	if root == nil {
		return
	}
	this.buildBST(root.Left)
	this.arr = append(this.arr, root.Val)
	this.buildBST(root.Right)
}

func (this *BSTIterator) Next() int {
	res := this.arr[0]
	this.arr = this.arr[1:]
	return res
}

func (this *BSTIterator) HasNext() bool {
	return len(this.arr) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	arr := []*TreeNode{root}
	left2Right := true

	for len(arr) > 0 {
		levelResult := make([]int, 0, len(arr))
		newArr := make([]*TreeNode, 0)

		for _, item := range arr {
			levelResult = append(levelResult, item.Val)

			if item.Left != nil {
				newArr = append(newArr, item.Left)
			}
			if item.Right != nil {
				newArr = append(newArr, item.Right)
			}
		}
		if !left2Right {
			reverseTreeNodes(levelResult)
		}
		arr = newArr
		result = append(result, levelResult)
		left2Right = !left2Right
	}

	return result
}

func reverseTreeNodes(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

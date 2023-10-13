package main

var res [][]int
var arr []int

func combine(n int, k int) [][]int {
	res = [][]int{}
	arr = []int{}
	combineHandle(1, n, k)
	return res
}
func combineHandle(cur, max, k int) {
	if len(arr) == k {
		tmp := make([]int, k)
		copy(tmp, arr)
		res = append(res, tmp)
	} else {
		for i := cur; i <= max; i++ {
			if max-cur+1 < k-len(arr) {
				break
			}
			arr = append(arr, i)
			combineHandle(i+1, max, k)
			arr = arr[:len(arr)-1]
		}
	}
}
